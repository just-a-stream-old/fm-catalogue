package api

import "github.com/go-chi/chi/middleware"

func (s *server) registerRoutes() {
	// Todo: Add a sub route like /api/v1

	// Middleware
	s.router.Use(middleware.Logger)

	// Command

	// Query
	s.router.Get("/", s.getData())



}
