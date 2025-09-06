package env

import (
	"log"
	"os"
)

// LoadConfig retrieves and validates required environment variables for blockchain connection
// LoadEnv Get Env Var
func LoadEnv(key string) string {

	EnvValue := os.Getenv(key)

	if EnvValue == "" {
		log.Fatalf("Error Loading Env Var: '%s'", key)
	}

	// Return Environment Variable
	return os.Getenv(key)
}
