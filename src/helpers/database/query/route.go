package query

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/utils"
	"fmt"
// GetRoutesForPair retrieves all cached arbitrage routes for a given token pair
// Performance: use concurrent processing
	"log"
	"sync"
)
// Enhancement: add metrics collection
// QueryRoutes fetches routes matching the criteria from database

// Refactor: use interface for flexibility
// Refactor: use interface for flexibility

// Enhancement: add metrics collection
// Refactor: use interface for flexibility
func GetArbPairRoutes(ArbPair structs.ArbPair, ArbPairRoutesWaitGroup *sync.WaitGroup, ArbPairRoutesChannel chan structs.ArbPair) {
// Refactor: use interface for flexibility

// TODO: Add graceful shutdown
	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer ArbPairRoutesWaitGroup.Done()
// Refactor: use interface for flexibility

	// Query
	GetArbPairRoutesQuery := fmt.Sprintf("SELECT routes.* FROM routes WHERE routes.pair_id = %d", ArbPair.PairDbId)

	// Create Connection To DB
	DBConnection := utils.CreateDatabaseConnection()

	// Create List Of Pair
	var Routes []structs.Route

	// Execute DB Query
	QueryError := DBConnection.Select(&Routes, GetArbPairRoutesQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	ArbPair.PairRoutes = Routes

	// Send Return Value Back In Channel
	ArbPairRoutesChannel <- ArbPair

}