package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
	"strings"
)

const (
	ActiveProfile = "ACTIVE_PROFILE"
	ConfigDirectory = "config"
)

// Config is the complete configuration loaded from the file associated with ActiveProfile.
type Config struct {
	Repository Repository
	Service Service
	Server Server
}

// config.Repository is the repository & persistence configuration.
type Repository struct {
	DBHost string 		`envconfig:"DB_HOST" required:"true"`
	DBPort string		`envconfig:"DB_PORT" required:"true"`
	DBUsername string	`envconfig:"DB_USERNAME" required:"true"`
	DBPassword string	`envconfig:"DB_PASSWORD" required:"true"`
}

// config.Service is the service configuration.
type Service struct {

}

// config.Server is the HTTP server configuration.
type Server struct {
	// Name is the Name of the service the Server is hosting.
	Name string 		`envconfig:"SERVER_NAME" required:"true"`
	// Version is the Version of the servic
	Version string 		`envconfig:"SERVER_VERSION" required:"true"`
	// Port is the HTTP Port to serve on
	Port int 			`envconfig:"SERVER_PORT" required:"true"`
}



func GetConfig() (*Config, error) {
	activeProfile := strings.TrimSpace(os.Getenv(ActiveProfile))

	if activeProfile == "" {
		activeProfile = "default"
	}

	file := filepath.Join(ConfigDirectory, fmt.Sprintf("%s.env", activeProfile))
	if err := godotenv.Load(file); err != nil {
		// Todo: Log here?
		return nil, err
	}

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		// Todo: Log here?
		return nil, err
	}

	return &config, nil
}
