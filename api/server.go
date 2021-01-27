package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

type server struct {
	httpServer http.Server
	service service.Service
	router *chi.Mux
	logger *zap.Logger
}
