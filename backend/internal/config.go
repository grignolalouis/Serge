package internal

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	DSN string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Database: DatabaseConfig{
			DSN: os.Getenv("DATABASE_URL"),
		},
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
			Host: getEnvOrDefault("SERVER_HOST", "localhost"),
		},
	}

	// Validate required configuration
	if cfg.Database.DSN == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	return cfg, nil
}

// getEnvOrDefault returns an environment variable value or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
