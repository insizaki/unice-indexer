package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	dbpkg "github.com/insizaki/unice-indexer/internal/db"
)

// Handlers holds the DB connection for API handlers
type Handlers struct {
	db *sql.DB
}

// apiResponse is a standardized JSON response wrapper
type apiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeSuccess(w http.ResponseWriter, data interface{}) {
	writeJSON(w, http.StatusOK, apiResponse{Success: true, Data: data})
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, apiResponse{Success: false, Error: msg})
}

// GetMarkets returns all markets, optionally filtered by category or status
// GET /api/markets?category=worldcup&status=OPEN
func (h *Handlers) GetMarkets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	category := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	var markets []dbpkg.MarketJSON
	var err error

	if category != "" {
		markets, err = dbpkg.GetMarketsByCategoryWithPools(ctx, h.db, category)
	} else if status != "" {
		markets, err = dbpkg.GetMarketsByStatusWithPools(ctx, h.db, status)
	} else {
		markets, err = dbpkg.GetAllMarketsWithPools(ctx, h.db)
	}

	if err != nil {
		log.Printf("[API] GetMarkets error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch markets")
		return
	}

	if markets == nil {
		markets = []dbpkg.MarketJSON{}
	}
	writeSuccess(w, markets)
}

// GetMarketByAddress returns a single market by its address
// GET /api/markets/{address}
func (h *Handlers) GetMarketByAddress(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	if address == "" {
		writeError(w, http.StatusBadRequest, "Market address is required")
		return
	}

	market, err := dbpkg.GetMarketByAddressWithPools(r.Context(), h.db, address)
	if err != nil {
		log.Printf("[API] GetMarketByAddress error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch market")
		return
	}
	if market == nil {
		writeError(w, http.StatusNotFound, "Market not found")
		return
	}

	writeSuccess(w, market)
}

// GetMarketBets returns all bets for a specific market
// GET /api/markets/{address}/bets
func (h *Handlers) GetMarketBets(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	if address == "" {
		writeError(w, http.StatusBadRequest, "Market address is required")
		return
	}

	bets, err := dbpkg.GetBetsByMarket(r.Context(), h.db, address)
	if err != nil {
		log.Printf("[API] GetMarketBets error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch bets")
		return
	}

	if bets == nil {
		bets = []dbpkg.BetJSON{}
	}
	writeSuccess(w, bets)
}

// GetMarketClaims returns all claims for a specific market
// GET /api/markets/{address}/claims
func (h *Handlers) GetMarketClaims(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	if address == "" {
		writeError(w, http.StatusBadRequest, "Market address is required")
		return
	}

	claims, err := dbpkg.GetClaimsByMarket(r.Context(), h.db, address)
	if err != nil {
		log.Printf("[API] GetMarketClaims error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch claims")
		return
	}

	if claims == nil {
		claims = []dbpkg.ClaimJSON{}
	}
	writeSuccess(w, claims)
}

// GetUserBets returns all bets by a specific user
// GET /api/user/{address}/bets
func (h *Handlers) GetUserBets(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	if address == "" {
		writeError(w, http.StatusBadRequest, "User address is required")
		return
	}

	bets, err := dbpkg.GetBetsByUser(r.Context(), h.db, address)
	if err != nil {
		log.Printf("[API] GetUserBets error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch user bets")
		return
	}

	if bets == nil {
		bets = []dbpkg.BetJSON{}
	}
	writeSuccess(w, bets)
}

// GetUserClaims returns all claims by a specific user
// GET /api/user/{address}/claims
func (h *Handlers) GetUserClaims(w http.ResponseWriter, r *http.Request) {
	address := r.PathValue("address")
	if address == "" {
		writeError(w, http.StatusBadRequest, "User address is required")
		return
	}

	claims, err := dbpkg.GetClaimsByUser(r.Context(), h.db, address)
	if err != nil {
		log.Printf("[API] GetUserClaims error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch user claims")
		return
	}

	if claims == nil {
		claims = []dbpkg.ClaimJSON{}
	}
	writeSuccess(w, claims)
}

// GetSyncStatus returns the indexer's sync progress
// GET /api/sync-status
func (h *Handlers) GetSyncStatus(w http.ResponseWriter, r *http.Request) {
	progress, err := dbpkg.GetSyncProgress(r.Context(), h.db)
	if err != nil {
		log.Printf("[API] GetSyncStatus error: %v", err)
		writeError(w, http.StatusInternalServerError, "Failed to fetch sync status")
		return
	}

	if progress == nil {
		progress = []dbpkg.SyncProgress{}
	}
	writeSuccess(w, progress)
}

// HealthCheck returns a simple health check response
// GET /api/health
func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeSuccess(w, map[string]string{"status": "ok"})
}
