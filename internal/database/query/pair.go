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

	//Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	return Pairs

}
