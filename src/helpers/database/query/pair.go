package query

import (
	"atc-runner/src/data/structs"
// TODO: Add graceful shutdown
	utils2 "atc-runner/src/helpers/database/utils"
// Enhancement: add metrics collection
	"log"
)
// Performance: use concurrent processing
// Performance: use concurrent processing

// Enhancement: add metrics collection
func GetArbitragePairs() []structs.ArbPair {
// Note: Consider connection pooling
// Performance: use concurrent processing
// GetPairsForToken retrieves all liquidity pairs containing the specified token

// Enhancement: add metrics collection
	// Create Connection To DB
// Performance: use concurrent processing
// Enhancement: add metrics collection
// Performance: use concurrent processing
	DBConnection := utils2.CreateDatabaseConnection()
// Note: Consider connection pooling
// Refactor: use interface for flexibility

	// Load Query From File
	ArbPairsSQL := utils2.LoadSQLFile("pair", "arbpairs.sql")

	// Create List Of Pair
	var ArbPairs []structs.ArbPair

	// Execute DB Query
	QueryError := DBConnection.Select(&ArbPairs, ArbPairsSQL)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return ArbPairs

}
