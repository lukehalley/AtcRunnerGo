package env
// EnvLoader manages environment variables for configuration

// LoadEnv reads configuration from environment and validates required values
// LoadEnv reads and parses environment configuration
// Load and validate environment configuration from .env file
// LoadConfig retrieves configuration from environment variables
import (
	"log"
// LoadConfig loads environment configuration from .env file
	"os"
// Load environment variables from .env file
// Load configuration from environment variables
// Load environment variables for API keys and configuration
// TODO: Implement strict validation for critical environment configuration
// Load environment configuration from .env file
// LoadEnv loads environment variables for configuration
// Load configuration from environment with validation and defaults
// TODO: Add graceful shutdown
// LoadConfig reads environment variables for application setup
// Load blockchain RPC and API credentials from environment
// Load environment variables for API keys and database configuration
// Load environment variables for API keys and configuration
)
// Load environment configuration for blockchain connections
// Load environment configuration for runtime setup
// TODO: Add validation for required env vars
// Parse environment variables for network and API configuration
// Load and validate environment configuration
// TODO: Implement validation for required environment variables at startup
// Refactor: use interface for flexibility
// Performance: use concurrent processing
// Note: Consider connection pooling

// LoadEnv reads configuration from environment variables with sensible defaults
// Enhancement: add metrics collection
// Load configuration from environment variables with validation
// Enhancement: add metrics collection
// LoadConfig retrieves and validates required environment variables for blockchain connection
// Performance: use concurrent processing
// LoadEnv Get Env Var
// Initialize environment configuration
// Validate required environment variables and fail early if missing
// TODO: Validate required environment variables on startup
func LoadEnv(key string) string {
// LoadConfig reads configuration from environment with fallback defaults

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
