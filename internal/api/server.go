package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"
)

// Server holds the HTTP server and shared DB connection
type Server struct {
	db         *sql.DB
	addr       string
	origins    []string
	httpServer *http.Server
}

// NewServer creates a new API server
func NewServer(db *sql.DB, addr string, corsOrigins string) *Server {
	origins := []string{"http://localhost:3000"}
	if corsOrigins != "" {
		origins = strings.Split(corsOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
		}
	}

	return &Server{
		db:      db,
		addr:    addr,
		origins: origins,
	}
}

// Start begins listening for HTTP requests (blocking)
func (s *Server) Start() {
	mux := http.NewServeMux()

	h := &Handlers{db: s.db}

	// Market endpoints
	mux.HandleFunc("GET /api/markets", h.GetMarkets)
	mux.HandleFunc("GET /api/markets/{address}", h.GetMarketByAddress)
	mux.HandleFunc("GET /api/markets/{address}/bets", h.GetMarketBets)
	mux.HandleFunc("GET /api/markets/{address}/claims", h.GetMarketClaims)

	// User endpoints
	mux.HandleFunc("GET /api/user/{address}/bets", h.GetUserBets)
	mux.HandleFunc("GET /api/user/{address}/claims", h.GetUserClaims)

	// Sync status
	mux.HandleFunc("GET /api/sync-status", h.GetSyncStatus)

	// Health check
	mux.HandleFunc("GET /api/health", h.HealthCheck)

	// Wrap with CORS middleware
	handler := s.corsMiddleware(mux)

	s.httpServer = &http.Server{
		Addr:         s.addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("[API] Server starting on %s", s.addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("[API] Server error: %v", err)
	}
}

// Shutdown gracefully shuts down the HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}
	log.Println("[API] Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}

// corsMiddleware adds CORS headers to responses
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowed := false
		for _, o := range s.origins {
			if o == "*" || o == origin {
				allowed = true
				break
			}
		}

		if allowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
