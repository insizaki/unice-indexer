package db

import (
	"context"
	"database/sql"
	"math/big"
)

type Market struct {
	Address            string
	Question           string
	Category           string
	BettingDeadline    uint64
	ResolutionDeadline uint64
	CollateralToken    string
	Status             string
	Outcome            int
	CreatedAtBlock     uint64
}

type Claim struct {
	TxHash        string
	LogIndex      int
	MarketAddress string
	UserAddress   string
	Amount        *big.Int
	ClaimType     string
	BlockNumber   uint64
}

type Bet struct {
	TxHash        string
	LogIndex      int
	MarketAddress string
	UserAddress   string
	IsYes         bool
	Amount        *big.Int
	BlockNumber   uint64
}

func UpdateMarketStatus(ctx context.Context, db *sql.DB, address string, status string) error {
	_, err := db.ExecContext(ctx,
		`UPDATE markets SET status = $2, updated_at = NOW() 
		WHERE address = $1`,
		address, status)
	return err
}

func SaveMarket(ctx context.Context, db *sql.DB, market *Market) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO markets (address, question, category, betting_deadline, resolution_deadline, collateral_token, status, outcome, created_at_block)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		 ON CONFLICT (address) DO NOTHING`,
		market.Address, market.Question, market.Category, market.BettingDeadline, market.ResolutionDeadline, market.CollateralToken, market.Status, market.Outcome, market.CreatedAtBlock)
	return err
}

func UpdateMarketOutcome(ctx context.Context, db *sql.DB, address string, outcome int) error {
	_, err := db.ExecContext(ctx,
		`UPDATE markets SET outcome = $2, updated_at = NOW() 
		WHERE address = $1`,
		address, outcome)
	return err
}

func SaveClaim(ctx context.Context, db *sql.DB, claim *Claim) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO claims (tx_hash, log_index, market_address, user_address, amount, claim_type, block_number)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 ON CONFLICT (tx_hash, log_index) DO NOTHING`,
		claim.TxHash, claim.LogIndex, claim.MarketAddress, claim.UserAddress, claim.Amount, claim.ClaimType, claim.BlockNumber)
	return err
}

func SaveBet(ctx context.Context, db *sql.DB, bet *Bet) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO bets (tx_hash, log_index, market_address, user_address, is_yes, amount, block_number)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)
		 ON CONFLICT (tx_hash, log_index) DO NOTHING`,
		bet.TxHash, bet.LogIndex, bet.MarketAddress, bet.UserAddress, bet.IsYes, bet.Amount, bet.BlockNumber)
	return err
}

func GetLastProcessedBlock(ctx context.Context, db *sql.DB, key string) (uint64, error) {
	var blockNumber uint64
	err := db.QueryRowContext(ctx, "SELECT last_processed_block FROM sync_progress WHERE key = $1", key).Scan(&blockNumber)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return blockNumber, err
}

func UpdateLastProcessedBlock(ctx context.Context, db *sql.DB, key string, blockNumber uint64) error {
	_, err := db.ExecContext(ctx,
		`INSERT INTO sync_progress (key, last_processed_block) 
		VALUES($1, $2) 
		ON CONFLICT (key) DO UPDATE SET last_processed_block = $2, updated_at = NOW()`,
		key, blockNumber)
	return err
}

func UpdateLastProcessedBlockTx(ctx context.Context, tx *sql.Tx, key string, blockNumber uint64) error {
	_, err := tx.ExecContext(ctx,
		`INSERT INTO sync_progress (key, last_processed_block) 
		VALUES($1, $2) 
		ON CONFLICT (key) DO UPDATE SET last_processed_block = $2, updated_at = NOW()`,
		key, blockNumber)
	return err
}

func GetAllMarketAddresses(ctx context.Context, db *sql.DB) ([]string, error) {
	rows, err := db.QueryContext(ctx, "SELECT address FROM markets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var addresses []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}
	return addresses, rows.Err()
}

func SaveMarketTx(ctx context.Context, tx *sql.Tx, m Market) error {
	_, err := tx.ExecContext(ctx, `
        INSERT INTO markets 
            (address, question, category, betting_deadline, resolution_deadline,
             status, outcome, created_at_block)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
        ON CONFLICT (address) DO NOTHING
    `, m.Address, m.Question, m.Category,
		m.BettingDeadline, m.ResolutionDeadline,
		m.Status, m.Outcome, m.CreatedAtBlock,
	)
	return err
}

func SaveBetTx(ctx context.Context, tx *sql.Tx, b Bet) error {
	_, err := tx.ExecContext(ctx, `
        INSERT INTO bets 
            (tx_hash, log_index, market_address, user_address, is_yes, amount, block_number)
        VALUES ($1,$2,$3,$4,$5,$6,$7)
        ON CONFLICT (tx_hash, log_index) DO NOTHING
    `, b.TxHash, b.LogIndex, b.MarketAddress,
		b.UserAddress, b.IsYes, b.Amount.String(), b.BlockNumber,
	)
	return err
}

func SaveClaimTx(ctx context.Context, tx *sql.Tx, c Claim) error {
	_, err := tx.ExecContext(ctx, `
        INSERT INTO claims 
            (tx_hash, log_index, market_address, user_address, amount, claim_type, block_number)
        VALUES ($1,$2,$3,$4,$5,$6,$7)
        ON CONFLICT (tx_hash, log_index) DO NOTHING
    `, c.TxHash, c.LogIndex, c.MarketAddress,
		c.UserAddress, c.Amount.String(), c.ClaimType, c.BlockNumber,
	)
	return err
}

func UpdateMarketStatusTx(ctx context.Context, tx *sql.Tx, address, status string, outcome int) error {
	_, err := tx.ExecContext(ctx, `
        UPDATE markets 
        SET status = $2, outcome = $3, updated_at = NOW()
        WHERE address = $1
    `, address, status, outcome)
	return err
}
