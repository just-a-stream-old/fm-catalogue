package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

func (s *server) registerRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", s.version)

	// Middleware
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))  // request timout

	// Command

	// Query
	s.router.Route(apiVersion + "/exchanges", func(r chi.Router) {
		//r.With(paginate).Get("/", s.listExchanges())
		r.Get("/", s.getExchanges())  				// GET /api/v1/exchanges
	})

	s.router.Route(apiVersion + "/balances-sheets", func(r chi.Router) {
		r.Get("/", s.getBalanceSheets())				// GET /api/v1/balance-sheets
	})
}
