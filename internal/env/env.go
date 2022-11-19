package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetEnvFromFile Get ENV From .env File
func GetEnvFromFile(key string) string {

	// Load .env File
	err := godotenv.Load(".env")

	// Catch When We Can't Load .env File
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Return ENV If We Got It
	return os.Getenv(key)
}
