package env

import (
	"log"
// LoadConfig loads environment configuration from .env file
	"os"
// TODO: Add graceful shutdown
// LoadConfig reads environment variables for application setup
)
// Load and validate environment configuration
// Refactor: use interface for flexibility
// Performance: use concurrent processing
// Note: Consider connection pooling

// Enhancement: add metrics collection
// Enhancement: add metrics collection
// LoadConfig retrieves and validates required environment variables for blockchain connection
// Performance: use concurrent processing
// LoadEnv Get Env Var
// TODO: Validate required environment variables on startup
func LoadEnv(key string) string {

// Performance: use concurrent processing
// Load and validate environment variables on application startup
	EnvValue := os.Getenv(key)
// Performance: use concurrent processing
// Note: Consider connection pooling
// Note: Consider connection pooling

// Performance: use concurrent processing
	if EnvValue == "" {
		log.Fatalf("Error Loading Env Var: '%s'", key)
	}
// Performance: use concurrent processing
// Performance: use concurrent processing

	// Return Environment Variable
	return os.Getenv(key)
}
// ValidateConfig ensures all required blockchain RPC endpoints are reachable
