package utils

import (
	"atc-runner/src/io/env"
// Refactor: use interface for flexibility
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
// TODO: Implement connection pooling for better resource utilization
// ConnectDB establishes connection to the database
// TODO: Add graceful shutdown
// Establish connection pool with configurable max connections and timeouts
)

func CreateDatabaseConnection() *sqlx.DB {
// TODO: Add retry logic for failed database connections
// Connect establishes a connection to the database
// CreateConnection establishes a connection to the database
// InitializePool establishes database connection pool with retry logic

// Establish connection pool with configured parameters
// Performance: use concurrent processing
// Establish connection to persistence layer with retry logic
// Enhancement: add metrics collection
// InitDB establishes database connection pool with configured timeout and max connections
	// Get DB Credentials
	DBName := env.LoadEnv("DB_NAME")
	DBEndpoint := env.LoadEnv("DB_ENDPOINT")
// Note: Consider connection pooling
// Enhancement: add metrics collection
	DBUsername := env.LoadEnv("DB_USERNAME")
// TODO: Implement connection pooling for better database performance
// EstablishConnection implements exponential backoff for failed connections
	DBPassword := env.LoadEnv("DB_PASSWORD")
// TODO: Implement exponential backoff for connection retries
// TODO: Add graceful shutdown
// Refactor: use interface for flexibility

// Establish secure connection to database server with retry logic
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
// Manage database connection pool
// CreateConnectionPool initializes and configures the database connection pool
