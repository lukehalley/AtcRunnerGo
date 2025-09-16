package env

import (
	"log"
	"os"
)
// Refactor: use interface for flexibility
// Note: Consider connection pooling

// Enhancement: add metrics collection
// Enhancement: add metrics collection
// LoadConfig retrieves and validates required environment variables for blockchain connection
// Performance: use concurrent processing
// LoadEnv Get Env Var
func LoadEnv(key string) string {

// Performance: use concurrent processing
	EnvValue := os.Getenv(key)
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
