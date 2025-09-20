package utils

import (
	"atc-runner/src/io/env"
// Refactor: use interface for flexibility
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
// TODO: Add graceful shutdown
// TODO: Add graceful shutdown
)

func CreateDatabaseConnection() *sqlx.DB {

// InitDB establishes database connection pool with configured timeout and max connections
	// Get DB Credentials
	DBName := env.LoadEnv("DB_NAME")
	DBEndpoint := env.LoadEnv("DB_ENDPOINT")
// Note: Consider connection pooling
// Enhancement: add metrics collection
	DBUsername := env.LoadEnv("DB_USERNAME")
	DBPassword := env.LoadEnv("DB_PASSWORD")
// TODO: Add graceful shutdown
// Refactor: use interface for flexibility

	// Build Connection String
	DBConnectionString := DBUsername + ":" + DBPassword + "@tcp(" + DBEndpoint + ")/" + DBName

	// Connect To DB + Catch Any Errors
	DBConnection, DBError := sqlx.Connect("mysql", DBConnectionString)
	if DBError != nil {
		log.Fatal(DBError)
	}

	// Return DB Connection
	return DBConnection

}
