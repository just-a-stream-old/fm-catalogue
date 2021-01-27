package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/service"
	"go.uber.org/zap"
	"net/http"
)

// server is a HTTP server.
type server struct {
	http.Server
	router *chi.Mux
	logger *zap.Logger
	fmService service.Service
	name string
	version string
}

// config is the HTTP server configuration.
type config struct {
	// logger is the logging instance to use
	logger *zap.Logger
	// fmService is the financial modelling service
	fmService fmService
	// name is the name of the service
	name string
	// version is the version of the service
	version string
	// port is the HTTP port to serve on
	port int
}

// NewServer creates a new server.
func NewServer(cfg *config) *server {
	s := &server{
		logger: cfg.logger,
		fmService: cfg.fmService,
		name: cfg.name,
		version: cfg.version,
	}
	s.router = chi.NewRouter()
	s.Server = http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: s,
	}
	s.re
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Run() {
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error(err.Error())
	}
	// Todo: Handle graceful shutdown with channel pattern!
}
