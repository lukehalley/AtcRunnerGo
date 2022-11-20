package query

import (
	"atc-runner/internal/database/structs"
	"atc-runner/internal/database/utils"
	. "github.com/ahmetalpbalkan/go-linq"
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

func GetArbitragePairs() ([]Group, int) {

	// Create Connection To DB
	DBConnection := utils.CreateDatabaseConnection()

	// Load Query From File
	ArbPairsSQL := utils.LoadSQLFile("pair", "arbpairsroutes.sql")

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

	// Group Arbitrage Pairs
	var GroupedArbitragePairs []Group
	From(ArbPairs).GroupByT(
		func(p structs.ArbPair) int { return p.RecipeGroupId },
		func(p structs.ArbPair) structs.ArbPair { return p },
	).ToSlice(&GroupedArbitragePairs)

	return GroupedArbitragePairs, len(ArbPairs)

}
