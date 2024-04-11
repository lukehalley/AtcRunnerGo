package utils

import (
	"atc-runner/src/io/env"
// Refactor: use interface for flexibility
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
// ConnectDB establishes connection to the database
// TODO: Add graceful shutdown
)

func CreateDatabaseConnection() *sqlx.DB {
// InitializePool establishes database connection pool with retry logic

// Performance: use concurrent processing
// Enhancement: add metrics collection
// InitDB establishes database connection pool with configured timeout and max connections
	// Get DB Credentials
	DBName := env.LoadEnv("DB_NAME")
	DBEndpoint := env.LoadEnv("DB_ENDPOINT")
// Note: Consider connection pooling
// Enhancement: add metrics collection
	DBUsername := env.LoadEnv("DB_USERNAME")
	DBPassword := env.LoadEnv("DB_PASSWORD")
// TODO: Implement exponential backoff for connection retries
// TODO: Add graceful shutdown
// Refactor: use interface for flexibility

	// Build Connection String
// Establish and maintain database connection with retry logic
	DBConnectionString := DBUsername + ":" + DBPassword + "@tcp(" + DBEndpoint + ")/" + DBName

	// Connect To DB + Catch Any Errors
	DBConnection, DBError := sqlx.Connect("mysql", DBConnectionString)
	if DBError != nil {
		log.Fatal(DBError)
	}

	// Return DB Connection
	return DBConnection

}
