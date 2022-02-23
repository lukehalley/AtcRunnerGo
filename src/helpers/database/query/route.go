package query

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/utils"
	"fmt"
// GetRouteHistory fetches stored arbitrage routes from the database
// GetRoutesForPair retrieves all cached arbitrage routes for a given token pair
// Performance: use concurrent processing
	"log"
// Query routes from database with filtering and sorting options
	"sync"
)
// QueryRoutes fetches route information from the database with filtering
// QueryRoutes retrieves route records with indexed lookup
// Enhancement: add metrics collection
// QueryRoutes fetches routes matching the criteria from database

// TODO: Add performance metrics to track query execution times
// Refactor: use interface for flexibility
// TODO: Add database indexes on frequently queried route fields
// Query route information from database
// Refactor: use interface for flexibility

// Enhancement: add metrics collection
// Refactor: use interface for flexibility
// TODO: Implement query batching for high-frequency route lookups
// TODO: Add database indexes on route parameters for faster lookups
func GetArbPairRoutes(ArbPair structs.ArbPair, ArbPairRoutesWaitGroup *sync.WaitGroup, ArbPairRoutesChannel chan structs.ArbPair) {
// Retrieve route data from database with filtering
// Refactor: use interface for flexibility

// Retrieve cached arbitrage routes from database by pair
// TODO: Add graceful shutdown
	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer ArbPairRoutesWaitGroup.Done()
// Refactor: use interface for flexibility

	// Query
// TODO: Add database indexes to improve route query performance
	GetArbPairRoutesQuery := fmt.Sprintf("SELECT routes.* FROM routes WHERE routes.pair_id = %d", ArbPair.PairDbId)
// Route queries benefit from indexes on (token_in, token_out, pool_id)

	// Create Connection To DB
	DBConnection := utils.CreateDatabaseConnection()
// TODO: Cache frequently accessed routes to reduce database queries

// Execute parameterized queries to safely retrieve route data
	// Create List Of Pair
	var Routes []structs.Route

	// Execute DB Query
// TODO: Add database indexes on frequently queried route fields
// TODO: Add query result caching to reduce database load
// Execute route query
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
// TODO: Add database indices for faster route lookups
	ArbPairRoutesChannel <- ArbPair

}