package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
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

// NewServer constructs a new server.
func NewServer(cfg *config.Server, logger *zap.Logger, fMService service.FMService) *server {
	s := &server{
		logger: logger,
		fMService: fMService,
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
	s.logger.Info(fmt.Sprintf("%s-%s running on port %s", s.name, s.version, s.Addr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error(err.Error())
	}
	// Todo: Handle graceful shutdown with channel pattern!
}
