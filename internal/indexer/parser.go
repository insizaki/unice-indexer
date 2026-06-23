package indexer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/insizaki/unice-indexer/internal/contracts"
	dbpkg "github.com/insizaki/unice-indexer/internal/db"
)

type Parser struct {
	idx        *Indexer
	factoryABI abi.ABI
	marketABI  abi.ABI
}

func NewParser(idx *Indexer) *Parser {
	fABI, err := contracts.UniceFactoryMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	mABI, err := contracts.UnicePredictionMarketMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	return &Parser{
		idx:        idx,
		factoryABI: *fABI,
		marketABI:  *mABI,
	}
}

func (p *Parser) HandleLog(ctx context.Context, tx *sql.Tx, l types.Log) error {
	if l.Address == p.idx.factoryAddr {
		return p.handleFactoryLog(ctx, tx, l)
	}
	return p.handleMarketLog(ctx, tx, l)
}

// --- Factory Events ---

func (p *Parser) handleFactoryLog(ctx context.Context, tx *sql.Tx, l types.Log) error {
	eventID := l.Topics[0]

	switch eventID {
	case p.factoryABI.Events["MarketCreated"].ID:
		return p.handleMarketCreated(ctx, tx, l)
	default:
		return nil // unknown event, skip
	}
}

func (p *Parser) handleMarketCreated(ctx context.Context, tx *sql.Tx, l types.Log) error {
	var event contracts.UniceFactoryMarketCreated
	if err := p.unpackFactoryEvent("MarketCreated", l, &event); err != nil {
		return err
	}

	market := dbpkg.Market{
		Address:            event.Market.Hex(),
		Question:           event.Question,
		Category:           event.Category,
		BettingDeadline:    event.BettingDeadline.Uint64(),
		ResolutionDeadline: event.ResolutionDeadline.Uint64(),
		Status:             "OPEN",
		Outcome:            0,
		CreatedAtBlock:     l.BlockNumber,
	}

	if err := dbpkg.SaveMarketTx(ctx, tx, market); err != nil {
		return fmt.Errorf("save market: %w", err)
	}

	// Dynamic watchlist update
	p.idx.addToWatchList(event.Market)
	return nil
}

// --- Market Events ---

func (p *Parser) handleMarketLog(ctx context.Context, tx *sql.Tx, l types.Log) error {
	eventID := l.Topics[0]

	switch eventID {
	case p.marketABI.Events["BetPlaced"].ID:
		return p.handleBetPlaced(ctx, tx, l)
	case p.marketABI.Events["ResultSubmitted"].ID:
		return p.handleResultSubmitted(ctx, tx, l)
	case p.marketABI.Events["ResultChallenged"].ID:
		return p.handleResultChallenged(ctx, tx, l)
	case p.marketABI.Events["MarketFinalized"].ID:
		return p.handleMarketFinalized(ctx, tx, l)
	case p.marketABI.Events["WinningsClaimed"].ID:
		return p.handleWinningsClaimed(ctx, tx, l)
	case p.marketABI.Events["ExpireRefundClaimed"].ID:
		return p.handleExpireRefundClaimed(ctx, tx, l)
	default:
		return nil
	}
}

func (p *Parser) handleBetPlaced(ctx context.Context, tx *sql.Tx, l types.Log) error {
	var event contracts.UnicePredictionMarketBetPlaced
	if err := p.unpackMarketEvent("BetPlaced", l, &event); err != nil {
		return err
	}

	bet := dbpkg.Bet{
		TxHash:        l.TxHash.Hex(),
		LogIndex:      int(l.Index),
		MarketAddress: l.Address.Hex(),
		UserAddress:   event.User.Hex(),
		IsYes:         event.IsYes,
		Amount:        event.Amount,
		BlockNumber:   l.BlockNumber,
	}

	return dbpkg.SaveBetTx(ctx, tx, bet)
}

func (p *Parser) handleResultSubmitted(ctx context.Context, tx *sql.Tx, l types.Log) error {
	var event contracts.UnicePredictionMarketResultSubmitted
	if err := p.unpackMarketEvent("ResultSubmitted", l, &event); err != nil {
		return err
	}

	outcome := int(event.Outcome) // 1 = YES, 2 = NO
	return dbpkg.UpdateMarketStatusTx(ctx, tx, l.Address.Hex(), "PENDING", outcome)
}

func (p *Parser) handleResultChallenged(ctx context.Context, tx *sql.Tx, l types.Log) error {
	return dbpkg.UpdateMarketStatusTx(ctx, tx, l.Address.Hex(), "OPEN", 0)
}

func (p *Parser) handleMarketFinalized(ctx context.Context, tx *sql.Tx, l types.Log) error {
	var event contracts.UnicePredictionMarketMarketFinalized
	if err := p.unpackMarketEvent("MarketFinalized", l, &event); err != nil {
		return err
	}

	outcome := int(event.Outcome)
	return dbpkg.UpdateMarketStatusTx(ctx, tx, l.Address.Hex(), "FINALIZED", outcome)
}

func (p *Parser) handleWinningsClaimed(ctx context.Context, tx *sql.Tx, l types.Log) error {
	var event contracts.UnicePredictionMarketWinningsClaimed
	if err := p.unpackMarketEvent("WinningsClaimed", l, &event); err != nil {
		return err
	}

	claim := dbpkg.Claim{
		TxHash:        l.TxHash.Hex(),
		LogIndex:      int(l.Index),
		MarketAddress: l.Address.Hex(),
		UserAddress:   event.User.Hex(),
		Amount:        event.Amount,
		ClaimType:     "WINNING",
		BlockNumber:   l.BlockNumber,
	}

	return dbpkg.SaveClaimTx(ctx, tx, claim)
}

func (p *Parser) handleExpireRefundClaimed(ctx context.Context, tx *sql.Tx, l types.Log) error {
	var event contracts.UnicePredictionMarketExpireRefundClaimed
	if err := p.unpackMarketEvent("ExpireRefundClaimed", l, &event); err != nil {
		return err
	}

	claim := dbpkg.Claim{
		TxHash:        l.TxHash.Hex(),
		LogIndex:      int(l.Index),
		MarketAddress: l.Address.Hex(),
		UserAddress:   event.User.Hex(),
		Amount:        event.Amount,
		ClaimType:     "REFUND",
		BlockNumber:   l.BlockNumber,
	}

	return dbpkg.SaveClaimTx(ctx, tx, claim)
}

// --- Unpack Helpers ---

func (p *Parser) unpackFactoryEvent(name string, l types.Log, out interface{}) error {
	// Non-indexed params ada di l.Data
	if len(l.Data) > 0 {
		if err := p.factoryABI.UnpackIntoInterface(out, name, l.Data); err != nil {
			return fmt.Errorf("unpack factory %s data: %w", name, err)
		}
	}
	// Indexed params ada di l.Topics
	if err := abi.ParseTopicsIntoMap(map[string]interface{}{}, p.factoryABI.Events[name].Inputs, l.Topics[1:]); err != nil {
		// handle indexed fields manual
	}
	// Market address adalah indexed topic[1]
	if name == "MarketCreated" && len(l.Topics) > 1 {
		type withMarket interface {
			SetMarket(common.Address)
		}
		if mc, ok := out.(*contracts.UniceFactoryMarketCreated); ok {
			mc.Market = common.HexToAddress(l.Topics[1].Hex())
		}
	}
	return nil
}

func (p *Parser) unpackMarketEvent(name string, l types.Log, out interface{}) error {
	if len(l.Data) > 0 {
		if err := p.marketABI.UnpackIntoInterface(out, name, l.Data); err != nil {
			return fmt.Errorf("unpack market %s: %w", name, err)
		}
	}
	return nil
}
