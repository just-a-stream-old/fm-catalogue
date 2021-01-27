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

// Config is the complete configuration loaded from the file associated with ActiveProfile
type Config struct {
	Server Server
}

// ServerConfig is configuration relating to the server
type Server struct {
	Port int `envconfig:"SERVER_PORT" required:"true"`
}

func GetConfig() (*Config, error) {
	activeProfile := strings.TrimSpace(os.Getenv(ActiveProfile))

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
