# Unice Indexer

> **EVM Event Indexer** for the Unice Prediction Market Protocol on **Arbitrum Sepolia**, written in **Go**.

The Unice Indexer is a blockchain event crawler that listens to on-chain events emitted by the Unice smart contracts (`UniceFactory` and `UnicePredictionMarket`) and persists them into a PostgreSQL database. It is designed for high reliability with checkpoint-based syncing, graceful shutdown, and dynamic market discovery.

---

## Table of Contents

- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Database Schema](#database-schema)
- [How It Works](#how-it-works)
- [Event Handling Reference](#event-handling-reference)
- [Development Guide](#development-guide)

---

## Architecture

```
┌──────────────────────────────────────────────────────────────────┐
│                        Unice Indexer                             │
│                                                                  │
│  ┌──────────┐    ┌────────────┐    ┌──────────┐    ┌──────────┐ │
│  │  main.go  │───▶│  Indexer   │───▶│  Parser  │───▶│   DB     │ │
│  │  (entry)  │    │  (engine)  │    │  (logs)  │    │ (store)  │ │
│  └──────────┘    └─────┬──────┘    └──────────┘    └──────────┘ │
│                        │                                         │
│                  ┌─────▼──────┐                                  │
│                  │  EthClient │                                   │
│                  │  (RPC)     │                                   │
│                  └────────────┘                                   │
└──────────────────────────────────────────────────────────────────┘
        │                                          │
        ▼                                          ▼
  Arbitrum Sepolia RPC                      PostgreSQL Database
```

**Flow Summary:**

1. The indexer connects to an Arbitrum Sepolia RPC node and a PostgreSQL database.
2. On startup, it loads the last processed block from the database (checkpoint).
3. It performs a **historical catch-up** — syncing all blocks from the checkpoint to the chain tip in configurable batches.
4. Once caught up, it switches to a **live polling** mode — checking for new blocks at a regular interval.
5. All retrieved event logs are decoded by the **Parser** and persisted via the **DB repository** layer within a single database transaction per batch.

---

## Project Structure

```
unice-indexer/
├── main.go                          # Application entry point
├── .env                             # Environment configuration
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
│
├── abis/                            # Contract ABI files
│   ├── UniceFactory.json            # Full artifact (ABI + bytecode)
│   ├── UniceFactory.abi.json        # ABI-only extract
│   ├── UnicePredictionMarket.json   # Full artifact (ABI + bytecode)
│   └── UnicePredictionMarket.abi.json
│
└── internal/
    ├── client/
    │   └── eth.go                   # Ethereum RPC client wrapper
    │
    ├── contracts/
    │   ├── factory.go               # Generated Go bindings (UniceFactory)
    │   └── market.go                # Generated Go bindings (UnicePredictionMarket)
    │
    ├── db/
    │   ├── db.go                    # PostgreSQL connection pool
    │   └── repository.go           # SQL queries — insert, update, select
    │
    └── indexer/
        ├── indexer.go               # Core indexer engine (catch-up, polling, lifecycle)
        └── parser.go                # Event log decoding and routing
```

### Package Responsibilities

| Package                | Description                                                                                  |
|------------------------|----------------------------------------------------------------------------------------------|
| `main`                 | Reads `.env`, initializes dependencies (DB, RPC), creates the `Indexer`, and starts it.      |
| `internal/client`      | Thin wrapper over `go-ethereum/ethclient` providing helper methods like `GetLatestBlockByNumber`. |
| `internal/contracts`   | Auto-generated Go bindings from ABIs via `abigen`. Contains ABI metadata and event structs.   |
| `internal/db`          | Database connection management and all SQL operations (CRUD for markets, bets, claims, sync). |
| `internal/indexer`     | Core business logic — block range iteration, log filtering, event parsing, and checkpointing.  |

---

## Prerequisites

- **Go** ≥ 1.21
- **PostgreSQL** ≥ 14
- **abigen** (from [go-ethereum tools](https://geth.ethereum.org/docs/tools/abigen)) — only needed to regenerate contract bindings
- An **Arbitrum Sepolia RPC endpoint** (public or private)

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/insizaki/unice-indexer.git
cd unice-indexer
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up the database

Create a PostgreSQL database and run the schema migrations:

```sql
-- Create the database
CREATE DATABASE unice;

-- Connect to the database and run schema
\c unice

-- Tables (see Database Schema section below for full DDL)
CREATE TABLE sync_progress (
    key VARCHAR(255) PRIMARY KEY,
    last_processed_block BIGINT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE markets (
    address VARCHAR(42) PRIMARY KEY,
    question TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
    betting_deadline BIGINT NOT NULL,
    resolution_deadline BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    outcome INT NOT NULL DEFAULT 0,
    created_at_block BIGINT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE bets (
    tx_hash VARCHAR(66) NOT NULL,
    log_index INT NOT NULL,
    market_address VARCHAR(42) NOT NULL REFERENCES markets(address),
    user_address VARCHAR(42) NOT NULL,
    is_yes BOOLEAN NOT NULL,
    amount NUMERIC(78, 0) NOT NULL,
    block_number BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (tx_hash, log_index)
);
CREATE INDEX idx_bets_market_address ON bets(market_address);
CREATE INDEX idx_bets_user_address ON bets(user_address);

CREATE TABLE claims (
    tx_hash VARCHAR(66) NOT NULL,
    log_index INT NOT NULL,
    market_address VARCHAR(42) NOT NULL REFERENCES markets(address),
    user_address VARCHAR(42) NOT NULL,
    amount NUMERIC(78, 0) NOT NULL,
    claim_type VARCHAR(20) NOT NULL,
    block_number BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (tx_hash, log_index)
);
```

### 4. Configure environment

Copy and edit the `.env` file:

```bash
cp .env.example .env
```

Update the values (see [Configuration](#configuration) for details).

### 5. Build and run

```bash
# Build
go build -o unice-indexer .

# Run
./unice-indexer
```

Or simply:

```bash
go run main.go
```

---

## Configuration

All configuration is managed via a `.env` file in the project root.

| Variable                 | Description                                                      | Default     |
|--------------------------|------------------------------------------------------------------|-------------|
| `RPC_URL`                | Arbitrum Sepolia JSON-RPC endpoint URL                           | *(required)* |
| `START_BLOCK`            | Block number to start indexing from (factory deployment block)   | `0`         |
| `FACTORY_ADDRESS`        | Deployed `UniceFactory` contract address                         | *(required)* |
| `DATABASE_URL`           | PostgreSQL connection string                                     | *(required)* |
| `BLOCK_BATCH_SIZE`       | Number of blocks to process per `FilterLogs` RPC call            | `2000`      |
| `CONFIRMATION_DEPTH`     | Number of blocks behind the chain tip to avoid reorg issues      | `3`         |
| `SYNC_INTERVAL_SECONDS`  | Polling interval (seconds) for live syncing after catch-up       | `3`         |

### Example `.env`

```env
RPC_URL=https://sepolia-rollup.arbitrum.io/rpc
START_BLOCK=45000000
FACTORY_ADDRESS=0xD66E625d40929123065ee95Bb67fFe08CCf8cFD1
DATABASE_URL=postgres://postgres:root@localhost:5432/unice?sslmode=disable
BLOCK_BATCH_SIZE=2000
CONFIRMATION_DEPTH=3
SYNC_INTERVAL_SECONDS=3
```

---

## Database Schema

### Entity Relationship Diagram

```
┌─────────────────┐       ┌──────────────────┐       ┌──────────────────┐
│  sync_progress   │       │     markets       │       │                  │
│─────────────────│       │──────────────────│◀──────│      bets        │
│ key (PK)         │       │ address (PK)      │       │──────────────────│
│ last_processed   │       │ question          │       │ tx_hash (PK)     │
│   _block         │       │ category          │       │ log_index (PK)   │
│ updated_at       │       │ betting_deadline  │       │ market_address   │
└─────────────────┘       │ resolution_       │       │ user_address     │
                           │   deadline        │       │ is_yes           │
                           │ status            │       │ amount           │
                           │ outcome           │       │ block_number     │
                           │ created_at_block  │       │ created_at       │
                           │ updated_at        │       └──────────────────┘
                           └──────────────────┘
                                    ▲
                                    │
                           ┌──────────────────┐
                           │     claims        │
                           │──────────────────│
                           │ tx_hash (PK)      │
                           │ log_index (PK)    │
                           │ market_address    │
                           │ user_address      │
                           │ amount            │
                           │ claim_type        │
                           │ block_number      │
                           │ created_at        │
                           └──────────────────┘
```

### Table Descriptions

| Table            | Purpose                                                          |
|------------------|------------------------------------------------------------------|
| `sync_progress`  | Tracks the last indexed block number per key (checkpoint).       |
| `markets`        | Stores metadata for each prediction market deployed by the factory. |
| `bets`           | Records every bet placed on any market.                          |
| `claims`         | Records all winnings claimed and expired refunds collected.      |

### Market Status Lifecycle

```
OPEN ──▶ PENDING ──▶ FINALIZED
  ▲         │
  └─────────┘  (if challenged)
```

| Status       | Meaning                                                  |
|--------------|----------------------------------------------------------|
| `OPEN`       | Market is active and accepting bets.                     |
| `PENDING`    | A result has been submitted but not yet finalized.       |
| `FINALIZED`  | The outcome is settled. Winners can claim payouts.       |

### Outcome Values

| Value | Meaning    |
|-------|------------|
| `0`   | Unresolved |
| `1`   | YES        |
| `2`   | NO         |
| `3`   | INVALID    |

---

## How It Works

### Startup Sequence

```
1. Load .env configuration
2. Connect to PostgreSQL
3. Connect to Arbitrum Sepolia RPC
4. Initialize Indexer with factory address
5. Load last processed block from DB (checkpoint)
   └── If no checkpoint → start from START_BLOCK
6. Load all known market addresses from DB → populate watchlist
7. Begin historical catch-up
8. Switch to live polling
```

### Historical Catch-Up

The indexer processes blocks in batches to avoid overwhelming the RPC node:

```
fromBlock = lastProcessedBlock
safeHead  = chainTip - CONFIRMATION_DEPTH

while fromBlock <= safeHead:
    toBlock = min(fromBlock + BATCH_SIZE - 1, safeHead)
    logs    = FilterLogs(fromBlock, toBlock, watchList)
    
    BEGIN TRANSACTION
        for each log:
            parse and store event data
        update sync_progress checkpoint
    COMMIT
    
    fromBlock = toBlock + 1
```

### Live Polling

After catching up, the indexer enters a ticker-based polling loop:

```
every SYNC_INTERVAL_SECONDS:
    lastBlock = getCheckpoint()
    safeHead  = chainTip - CONFIRMATION_DEPTH
    
    if lastBlock >= safeHead:
        continue  (already caught up)
    
    processRange(lastBlock + 1, safeHead)
```

### Dynamic Watchlist

The indexer maintains an in-memory list of contract addresses to monitor:

- **On startup:** Initialized with the `FACTORY_ADDRESS` + all known market addresses from the database.
- **On `MarketCreated` event:** The new market address is appended to the watchlist in real-time, so subsequent block batches will include logs from the new market.

### Confirmation Depth

To avoid indexing logs from blocks that may be reorganized, the indexer always stays `CONFIRMATION_DEPTH` blocks behind the chain tip. For example, with `CONFIRMATION_DEPTH=3` and a chain tip at block 1000, the indexer will only process up to block 997.

### Graceful Shutdown

The indexer listens for `SIGINT` and `SIGTERM` signals. Upon receiving a shutdown signal:

1. The context is cancelled, stopping the polling loop.
2. Any in-progress database transaction is rolled back (no partial writes).
3. The last successfully committed checkpoint is preserved.
4. Database and RPC connections are closed.

---

## Event Handling Reference

### Factory Events

| Event           | Source Contract  | Handler                    | Action                                                    |
|-----------------|------------------|----------------------------|-----------------------------------------------------------|
| `MarketCreated` | `UniceFactory`   | `handleMarketCreated`      | Insert new market into `markets` table, add to watchlist  |

### Market Events

| Event                 | Source Contract           | Handler                        | Action                                                      |
|-----------------------|---------------------------|--------------------------------|-------------------------------------------------------------|
| `BetPlaced`           | `UnicePredictionMarket`   | `handleBetPlaced`              | Insert record into `bets` table                             |
| `ResultSubmitted`     | `UnicePredictionMarket`   | `handleResultSubmitted`        | Update market status → `PENDING`, set outcome               |
| `ResultChallenged`    | `UnicePredictionMarket`   | `handleResultChallenged`       | Update market status → `OPEN`, reset outcome to `0`         |
| `MarketFinalized`     | `UnicePredictionMarket`   | `handleMarketFinalized`        | Update market status → `FINALIZED`, set final outcome       |
| `WinningsClaimed`     | `UnicePredictionMarket`   | `handleWinningsClaimed`        | Insert record into `claims` with `claim_type = 'WINNING'`   |
| `ExpireRefundClaimed` | `UnicePredictionMarket`   | `handleExpireRefundClaimed`    | Insert record into `claims` with `claim_type = 'REFUND'`    |

### Event Parsing Strategy

- **Non-indexed parameters** are decoded from `log.Data` using `abi.UnpackIntoInterface`.
- **Indexed parameters** (e.g., the market address in `MarketCreated`) are extracted manually from `log.Topics`.
- All inserts use `ON CONFLICT ... DO NOTHING` to ensure idempotency — re-processing the same block range will not create duplicate records.

---

## Development Guide

### Regenerating Contract Bindings

If the contract ABIs are updated, regenerate the Go bindings:

```bash
# Install abigen (if not already installed)
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Generate factory bindings
abigen --abi=abis/UniceFactory.abi.json \
       --pkg=contracts \
       --out=internal/contracts/factory.go

# Generate market bindings
abigen --abi=abis/UnicePredictionMarket.abi.json \
       --pkg=contracts \
       --out=internal/contracts/market.go
```

### Building

```bash
go build -o unice-indexer .
```

### Key Dependencies

| Dependency                      | Purpose                                     |
|---------------------------------|---------------------------------------------|
| `github.com/ethereum/go-ethereum` | Ethereum client, ABI parsing, log filtering |
| `github.com/lib/pq`             | PostgreSQL driver for `database/sql`         |
| `github.com/joho/godotenv`      | Load environment variables from `.env` file  |

### Adding a New Event

To index a new event emitted by the smart contracts:

1. **Update the ABI** — Add the new ABI JSON file to `abis/` and regenerate bindings.
2. **Add handler** — In `parser.go`, add a new `case` in the `handleMarketLog` or `handleFactoryLog` switch statement.
3. **Add DB model** — If the event requires a new table, add the struct in `repository.go` and create the corresponding `Save*Tx` function.
4. **Create migration** — Add the DDL for the new table to your migration scripts.

### Troubleshooting

| Problem                                | Solution                                                                  |
|----------------------------------------|---------------------------------------------------------------------------|
| `pq: column "X" does not exist`        | Column name in SQL query doesn't match the database schema. Verify DDL.   |
| `Error: filter logs: ...`              | RPC node may be rate-limiting. Reduce `BLOCK_BATCH_SIZE`.                 |
| Indexer re-processes old blocks        | Check that `sync_progress` table has the correct checkpoint value.        |
| `panic: ... GetAbi()`                 | ABI JSON may be malformed. Re-generate bindings with `abigen`.            |
| Duplicate entries                      | Not possible — all inserts use `ON CONFLICT DO NOTHING`.                  |

---

## License

This project is part of the Unice Protocol.
