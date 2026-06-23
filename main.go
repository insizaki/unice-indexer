package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/insizaki/unice-indexer/internal/client"
	dbpkg "github.com/insizaki/unice-indexer/internal/db"
	"github.com/insizaki/unice-indexer/internal/indexer"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init DB
	db, err := dbpkg.NewConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	// Init Eth client
	ethClient, err := client.NewEthClient(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("ETH client failed: %v", err)
	}
	defer ethClient.Close()

	// Load config dari .env
	startBlock := envUint64("START_BLOCK", 0)
	batchSize := envUint64("BLOCK_BATCH_SIZE", 2000)
	confirmDepth := envUint64("CONFIRMATION_DEPTH", 3)
	syncSecs := envUint64("SYNC_INTERVAL_SECONDS", 3)

	// Init indexer
	idx, err := indexer.NewIndexer(
		ethClient.Client,
		db,
		os.Getenv("FACTORY_ADDRESS"),
		startBlock,
		batchSize,
		confirmDepth,
		time.Duration(syncSecs)*time.Second,
	)
	if err != nil {
		log.Fatalf("Failed to create indexer: %v", err)
	}

	// Run
	indexer.Run(context.Background(), idx)
}

func envUint64(key string, fallback uint64) uint64 {
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
