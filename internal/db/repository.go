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

// ─── Read-Only Query Functions (for REST API) ───

type SyncProgress struct {
	Key                string `json:"key"`
	LastProcessedBlock uint64 `json:"last_processed_block"`
}

// MarketJSON is a JSON-friendly version of Market with computed pool amounts
type MarketJSON struct {
	Address            string `json:"address"`
	Question           string `json:"question"`
	Category           string `json:"category"`
	BettingDeadline    uint64 `json:"betting_deadline"`
	ResolutionDeadline uint64 `json:"resolution_deadline"`
	Status             string `json:"status"`
	Outcome            int    `json:"outcome"`
	CreatedAtBlock     uint64 `json:"created_at_block"`
	TotalYesPool       string `json:"total_yes_pool"`
	TotalNoPool        string `json:"total_no_pool"`
	YesOdds            int    `json:"yes_odds"`
	NoOdds             int    `json:"no_odds"`
}

// BetJSON is a JSON-friendly version of Bet
type BetJSON struct {
	TxHash        string `json:"tx_hash"`
	LogIndex      int    `json:"log_index"`
	MarketAddress string `json:"market_address"`
	UserAddress   string `json:"user_address"`
	IsYes         bool   `json:"is_yes"`
	Amount        string `json:"amount"`
	BlockNumber   uint64 `json:"block_number"`
}

// ClaimJSON is a JSON-friendly version of Claim
type ClaimJSON struct {
	TxHash        string `json:"tx_hash"`
	LogIndex      int    `json:"log_index"`
	MarketAddress string `json:"market_address"`
	UserAddress   string `json:"user_address"`
	Amount        string `json:"amount"`
	ClaimType     string `json:"claim_type"`
	BlockNumber   uint64 `json:"block_number"`
}

func GetAllMarketsWithPools(ctx context.Context, db *sql.DB) ([]MarketJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT m.address, m.question, m.category, m.betting_deadline, m.resolution_deadline,
		       m.status, m.outcome, m.created_at_block,
		       COALESCE(SUM(CASE WHEN b.is_yes = true THEN b.amount ELSE 0 END), 0) AS total_yes_pool,
		       COALESCE(SUM(CASE WHEN b.is_yes = false THEN b.amount ELSE 0 END), 0) AS total_no_pool
		FROM markets m
		LEFT JOIN bets b ON b.market_address = m.address
		GROUP BY m.address
		ORDER BY m.created_at_block DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanMarketsWithPools(rows)
}

func GetMarketsByCategoryWithPools(ctx context.Context, db *sql.DB, category string) ([]MarketJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT m.address, m.question, m.category, m.betting_deadline, m.resolution_deadline,
		       m.status, m.outcome, m.created_at_block,
		       COALESCE(SUM(CASE WHEN b.is_yes = true THEN b.amount ELSE 0 END), 0) AS total_yes_pool,
		       COALESCE(SUM(CASE WHEN b.is_yes = false THEN b.amount ELSE 0 END), 0) AS total_no_pool
		FROM markets m
		LEFT JOIN bets b ON b.market_address = m.address
		WHERE m.category = $1
		GROUP BY m.address
		ORDER BY m.created_at_block DESC
	`, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanMarketsWithPools(rows)
}

func GetMarketsByStatusWithPools(ctx context.Context, db *sql.DB, status string) ([]MarketJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT m.address, m.question, m.category, m.betting_deadline, m.resolution_deadline,
		       m.status, m.outcome, m.created_at_block,
		       COALESCE(SUM(CASE WHEN b.is_yes = true THEN b.amount ELSE 0 END), 0) AS total_yes_pool,
		       COALESCE(SUM(CASE WHEN b.is_yes = false THEN b.amount ELSE 0 END), 0) AS total_no_pool
		FROM markets m
		LEFT JOIN bets b ON b.market_address = m.address
		WHERE m.status = $1
		GROUP BY m.address
		ORDER BY m.created_at_block DESC
	`, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanMarketsWithPools(rows)
}

func GetMarketByAddressWithPools(ctx context.Context, db *sql.DB, address string) (*MarketJSON, error) {
	row := db.QueryRowContext(ctx, `
		SELECT m.address, m.question, m.category, m.betting_deadline, m.resolution_deadline,
		       m.status, m.outcome, m.created_at_block,
		       COALESCE(SUM(CASE WHEN b.is_yes = true THEN b.amount ELSE 0 END), 0) AS total_yes_pool,
		       COALESCE(SUM(CASE WHEN b.is_yes = false THEN b.amount ELSE 0 END), 0) AS total_no_pool
		FROM markets m
		LEFT JOIN bets b ON b.market_address = m.address
		WHERE m.address = $1
		GROUP BY m.address
	`, address)

	var m MarketJSON
	var yesPool, noPool string
	err := row.Scan(&m.Address, &m.Question, &m.Category, &m.BettingDeadline,
		&m.ResolutionDeadline, &m.Status, &m.Outcome, &m.CreatedAtBlock,
		&yesPool, &noPool)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	m.TotalYesPool = yesPool
	m.TotalNoPool = noPool
	computeOdds(&m)
	return &m, nil
}

func GetBetsByMarket(ctx context.Context, db *sql.DB, marketAddress string) ([]BetJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT tx_hash, log_index, market_address, user_address, is_yes, amount, block_number
		FROM bets WHERE market_address = $1
		ORDER BY block_number DESC
	`, marketAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanBets(rows)
}

func GetBetsByUser(ctx context.Context, db *sql.DB, userAddress string) ([]BetJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT tx_hash, log_index, market_address, user_address, is_yes, amount, block_number
		FROM bets WHERE user_address = $1
		ORDER BY block_number DESC
	`, userAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanBets(rows)
}

func GetClaimsByMarket(ctx context.Context, db *sql.DB, marketAddress string) ([]ClaimJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT tx_hash, log_index, market_address, user_address, amount, claim_type, block_number
		FROM claims WHERE market_address = $1
		ORDER BY block_number DESC
	`, marketAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanClaims(rows)
}

func GetClaimsByUser(ctx context.Context, db *sql.DB, userAddress string) ([]ClaimJSON, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT tx_hash, log_index, market_address, user_address, amount, claim_type, block_number
		FROM claims WHERE user_address = $1
		ORDER BY block_number DESC
	`, userAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanClaims(rows)
}

func GetSyncProgress(ctx context.Context, db *sql.DB) ([]SyncProgress, error) {
	rows, err := db.QueryContext(ctx, `SELECT key, last_processed_block FROM sync_progress`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []SyncProgress
	for rows.Next() {
		var sp SyncProgress
		if err := rows.Scan(&sp.Key, &sp.LastProcessedBlock); err != nil {
			return nil, err
		}
		results = append(results, sp)
	}
	return results, rows.Err()
}

// ─── Scan Helpers ───

func scanMarketsWithPools(rows *sql.Rows) ([]MarketJSON, error) {
	var markets []MarketJSON
	for rows.Next() {
		var m MarketJSON
		var yesPool, noPool string
		if err := rows.Scan(&m.Address, &m.Question, &m.Category, &m.BettingDeadline,
			&m.ResolutionDeadline, &m.Status, &m.Outcome, &m.CreatedAtBlock,
			&yesPool, &noPool); err != nil {
			return nil, err
		}
		m.TotalYesPool = yesPool
		m.TotalNoPool = noPool
		computeOdds(&m)
		markets = append(markets, m)
	}
	return markets, rows.Err()
}

func computeOdds(m *MarketJSON) {
	// Parse pool amounts (stored as wei strings, convert to float for odds calc)
	yes := new(big.Int)
	no := new(big.Int)
	yes.SetString(m.TotalYesPool, 10)
	no.SetString(m.TotalNoPool, 10)

	total := new(big.Int).Add(yes, no)
	if total.Sign() == 0 {
		m.YesOdds = 50
		m.NoOdds = 50
		return
	}

	// Odds as percentage (basis-100)
	hundred := big.NewInt(100)
	yesOdds := new(big.Int).Mul(yes, hundred)
	yesOdds.Div(yesOdds, total)
	m.YesOdds = int(yesOdds.Int64())
	m.NoOdds = 100 - m.YesOdds
}

func scanBets(rows *sql.Rows) ([]BetJSON, error) {
	var bets []BetJSON
	for rows.Next() {
		var b BetJSON
		if err := rows.Scan(&b.TxHash, &b.LogIndex, &b.MarketAddress,
			&b.UserAddress, &b.IsYes, &b.Amount, &b.BlockNumber); err != nil {
			return nil, err
		}
		bets = append(bets, b)
	}
	return bets, rows.Err()
}

func scanClaims(rows *sql.Rows) ([]ClaimJSON, error) {
	var claims []ClaimJSON
	for rows.Next() {
		var c ClaimJSON
		if err := rows.Scan(&c.TxHash, &c.LogIndex, &c.MarketAddress,
			&c.UserAddress, &c.Amount, &c.ClaimType, &c.BlockNumber); err != nil {
			return nil, err
		}
		claims = append(claims, c)
	}
	return claims, rows.Err()
}
