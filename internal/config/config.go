package config

import (
	"fmt"
	"os"
)

const (
	// EnvAPIKey is the environment variable name for the Gemini API key
	EnvAPIKey = "GEMINI_API_KEY"
)

// Config holds the application configuration
type Config struct {
	APIKey string
}

// Load loads and validates the application configuration from environment variables
func Load() (*Config, error) {
	apiKey := os.Getenv(EnvAPIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("%s environment variable is not set. Please set it in your .env file or export it in your shell", EnvAPIKey)
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}
