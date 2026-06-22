package indexer

import (
	"database/sql"
	"go/parser"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/insizaki/unice-indexer/internal/db"
)

type Indexer struct {
	client       *ethclient.Client
	db           *db.Client
	parser       *parser.Parser
	watchList    []common.AddressLength
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
