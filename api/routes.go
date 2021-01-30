package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

func (s *server) registerRoutes() {
	// Todo: Add a sub route like /api/v1

	// Middleware
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))  // request timout

	// Command

	// Query
	s.router.Route("/api/v1/exchanges", func(r chi.Router) {
		//r.With(paginate).Get("/", s.listExchanges())
		r.Get("/", s.listExchanges())
	})
}
