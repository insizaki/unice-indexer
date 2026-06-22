package db

import (
	"context"
	"database/sql"
	"math/big"
)

type Market struct {
	Address            int64
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
		`INSERT INTO markets (address, question, category, betting_deadline, resolution_deadline, collateral_token, status, outcome, created_at_block, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		 ON CONFLICT (address) DO NOTHING`,
		//  SET question = EXCLUDED.question,
		//  category = EXCLUDED.category,
		//  betting_deadline = EXCLUDED.betting_deadline,
		//  resolution_deadline = EXCLUDED.resolution_deadline,
		//  collateral_token = EXCLUDED.collateral_token,
		//  status = EXCLUDED.status,
		//  outcome = EXCLUDED.outcome,
		//  created_at_block = EXCLUDED.created_at_block,
		//  updated_at = NOW()`,
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
