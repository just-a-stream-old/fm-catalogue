package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/service"
	"go.uber.org/zap"
	"net/http"
)

type apiError struct {
	Error 	string	`json:"Error"`
	Message string	`json:"Message"`
}

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
func NewServer(cfg *config.Server, logger *zap.Logger, fMService service.FMService) (*server, error) {
	if cfg.Name == "" {
		return nil, fmt.Errorf("server name is required to construct the server, and is empty")
	}
	if cfg.Version == "" {
		return nil, fmt.Errorf("server version is required to construct the server, and is empty")
	}

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

	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Run() {
	// Todo: Handle graceful shutdown with channel pattern!
	s.logger.Info(fmt.Sprintf("%s-%s running on port %s", s.name, s.version, s.Addr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error(err.Error())
	}
}

func (s *server) respondError(w http.ResponseWriter, r *http.Request, code int, err error, message string) {
	s.logger.Error(err.Error())
	s.respondJSON(w, code, apiError{
		Error: err.Error(),
		Message: message,
	})
}

func (s *server) respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
