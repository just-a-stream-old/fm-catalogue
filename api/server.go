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
	fMService service.FMService
	name string
	version string
}

// Config is the HTTP server configuration.
type Config struct {
	// Logger is the logging instance to use
	Logger *zap.Logger
	// service is the financial modelling service
	FMService service.FMService
	// Name is the Name of the service
	Name string
	// Version is the Version of the service
	Version string
	// Port is the HTTP Port to serve on
	Port int
}

// NewServer creates a new server.
func NewServer(cfg *Config) *server {
	s := &server{
		logger: cfg.Logger,
		fMService: cfg.fmService,
		name: cfg.Name,
		version: cfg.Version,
	}
	s.router = chi.NewRouter()
	s.Server = http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Port),
		Handler: s,
	}
	s.registerRoutes()
	return s
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
