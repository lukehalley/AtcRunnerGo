package query

import (
	"atc-runner/internal/database/structs"
	"atc-runner/internal/database/utils"
	"log"
)

func GetAllPairs() []structs.Pair {

	// Query
	GetAllPairsQuery := "SELECT pairs.* FROM pairs LIMIT 100"

	// Create Connection To DB
	DBConnection := utils.CreateDatabaseConnection()

	// Create List Of Pair
	var Pairs []structs.Pair

	// Execute DB Query
	QueryError := DBConnection.Select(&Pairs, GetAllPairsQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Pairs

}

func GetArbitragePairs() []structs.Pair {

	// Create Connection To DB
	DBConnection := utils.CreateDatabaseConnection()

	// Load Query From File
	ArbPairsSQL := utils.LoadSQLFile("pair", "arbpairs.sql")

	// Create List Of Pair
	var Pairs []structs.Pair

	// Execute DB Query
	QueryError := DBConnection.Select(&Pairs, ArbPairsSQL)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Pairs

}
