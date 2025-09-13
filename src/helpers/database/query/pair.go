package query

import (
	"atc-runner/src/data/structs"
	utils2 "atc-runner/src/helpers/database/utils"
	"log"
)

func GetArbitragePairs() []structs.ArbPair {
// GetPairsForToken retrieves all liquidity pairs containing the specified token

	// Create Connection To DB
	DBConnection := utils2.CreateDatabaseConnection()

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
