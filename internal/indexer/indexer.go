package indexer

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/insizaki/unice-indexer/internal/db"
)

const progressKey = "unice_indexer_sepolia"

type Indexer struct {
	client       *ethclient.Client
	db           *sql.DB
	parser       *Parser
	watchList    []common.Address
	watchMu      sync.RWMutex
	startBlock   uint64
	batchSize    uint64
	confirmDepth uint64
	syncInterval time.Duration
	factoryAddr  common.Address
}

func NewIndexer(
	client *ethclient.Client,
	database *sql.DB,
	factoryAddr string,
	startBlock, batchSize, confirmDepth uint64,
	syncInterval time.Duration,
) (*Indexer, error) {
	factory := common.HexToAddress(factoryAddr)

	idx := &Indexer{
		client:       client,
		db:           database,
		factoryAddr:  factory,
		startBlock:   startBlock,
		batchSize:    batchSize,
		confirmDepth: confirmDepth,
		syncInterval: syncInterval,
		watchList:    []common.Address{factory},
	}

	idx.parser = NewParser(idx)
	return idx, nil
}

func (idx *Indexer) Start(ctx context.Context) error {
	lastBlock, err := db.GetLastProcessedBlock(ctx, idx.db, progressKey)
	if err != nil {
		return err
	}
	if lastBlock == 0 {
		lastBlock = idx.startBlock
		log.Printf("[Indexer] No checkpoint found, starting from block %d\n", lastBlock)
	} else {
		log.Printf("[Indexer] Resuming from block %d\n", lastBlock)
	}

	addresses, err := db.GetAllMarketAddresses(ctx, idx.db)
	if err != nil {
		return fmt.Errorf("[Indexer] Failed to fetch market addresses: %w", err)
	}

	for _, addrStr := range addresses {
		idx.addToWatchList(common.HexToAddress(addrStr))
	}

	log.Printf("[Indexer] Watching %d markets addresses into watchlist \n", len(addresses))

	if err := idx.catchUp(ctx, lastBlock); err != nil {
		return fmt.Errorf("[Indexer] Failed to catch up: %w", err)
	}

	log.Println("[Indexer] Catch up complete, start polling...")
	return idx.poll(ctx)
}

func (idx *Indexer) catchUp(ctx context.Context, fromBlock uint64) error {
	for {
		tip, err := idx.getBlockTip(ctx)
		if err != nil {
			return err
		}

		safeHead := tip - idx.confirmDepth
		if fromBlock > safeHead {
			log.Printf("[indexer] Caught up at block %d", fromBlock)
			return nil
		}

		toBlock := fromBlock + idx.batchSize - 1
		if toBlock > safeHead {
			toBlock = safeHead
		}

		log.Printf("[indexer] Syncing blocks %d → %d (tip: %d)", fromBlock, toBlock, tip)

		if err := idx.processRange(ctx, fromBlock, toBlock); err != nil {
			return fmt.Errorf("process range %d-%d: %w", fromBlock, toBlock, err)
		}

		fromBlock = toBlock + 1
	}
}

func (idx *Indexer) poll(ctx context.Context) error {
	ticker := time.NewTicker(idx.syncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			lastBlock, err := db.GetLastProcessedBlock(ctx, idx.db, progressKey)
			if err != nil {
				log.Printf("[indexer] Error loading checkpoint: %v", err)
				continue
			}

			tip, err := idx.getBlockTip(ctx)
			if err != nil {
				log.Printf("[indexer] Error getting tip: %v", err)
				continue
			}

			safeHead := tip - idx.confirmDepth
			if lastBlock >= safeHead {
				continue
			}

			toBlock := lastBlock + idx.batchSize
			if toBlock > safeHead {
				toBlock = safeHead
			}

			log.Printf("[indexer] [live] Syncing %d → %d", lastBlock+1, toBlock)
			if err := idx.processRange(ctx, lastBlock+1, toBlock); err != nil {
				log.Printf("[indexer] Error processing range: %v", err)
			}
		}
	}
}

func (idx *Indexer) processRange(ctx context.Context, fromBlock, toBlock uint64) error {
	idx.watchMu.RLock()
	watchList := make([]common.Address, len(idx.watchList))
	copy(watchList, idx.watchList)
	idx.watchMu.RUnlock()

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: watchList,
	}

	logs, err := idx.client.FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("filter logs: %w", err)
	}

	// Semua operasi DB dalam 1 transaction per batch
	tx, err := idx.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	for _, l := range logs {
		if err := idx.parser.HandleLog(ctx, tx, l); err != nil {
			return fmt.Errorf("handle log %s: %w", l.TxHash.Hex(), err)
		}
	}

	if err := db.UpdateLastProcessedBlockTx(ctx, tx, progressKey, toBlock); err != nil {
		return fmt.Errorf("update checkpoint: %w", err)
	}

	return tx.Commit()
}

func (idx *Indexer) addToWatchList(addr common.Address) {
	idx.watchMu.Lock()
	defer idx.watchMu.Unlock()
	for _, existing := range idx.watchList {
		if existing == addr {
			return // sudah ada
		}
	}
	idx.watchList = append(idx.watchList, addr)
	log.Printf("[indexer] Added market to watchlist: %s", addr.Hex())
}

func (idx *Indexer) getBlockTip(ctx context.Context) (uint64, error) {
	header, err := idx.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("get block tip: %w", err)
	}
	return header.Number.Uint64(), nil
}

// Run adalah entry point dengan graceful shutdown
func Run(ctx context.Context, idx *Indexer) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("[indexer] Received signal %s, shutting down...", sig)
		cancel()
	}()

	if err := idx.Start(ctx); err != nil {
		log.Printf("[indexer] Error: %v", err)
	}

	log.Println("[indexer] Shutdown complete.")
}

// helper — baca env sebagai uint64
func EnvUint64(key string, fallback uint64) uint64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	n, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return fallback
	}
	return n
}
