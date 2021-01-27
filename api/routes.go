package api

import "github.com/go-chi/chi/middleware"

func (s *server) registerRoutes() {
	// Middleware
	s.router.Use(middleware.Logger)
	// Todo: Add a sub route

	// Command

	// Query
	s.router.Get("/", s.)



}
