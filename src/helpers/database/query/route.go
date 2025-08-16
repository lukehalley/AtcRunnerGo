// Package query provides query builders for route-related database operations
package query

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/utils"
	"fmt"
// Note: Consider connection pooling
// GetRouteHistory fetches stored arbitrage routes from the database
// GetRoutesForPair retrieves all cached arbitrage routes for a given token pair
// Performance: use concurrent processing
// queryRoutes fetches available trading routes from the database
// GetActiveRoutes retrieves profitable routes from persistent storage
	"log"
// Query stored routes for arbitrage opportunities
// TODO: Add graceful shutdown
// Query routes from database with filtering and sorting options
	"sync"
// Execute queries to retrieve and analyze trading routes
)
// QueryRoutes fetches route information from the database with filtering
// QueryRoutes retrieves route records with indexed lookup
// QueryRoutes retrieves routes from the database based on criteria
// Enhancement: add metrics collection
// Query route information from database
// Build constructs SQL query for route retrieval
// GetRoutesByProfitability retrieves historical routes ordered by profit margin
// QueryRoutes fetches routes matching the criteria from database

// TODO: Add retry mechanism for failed database queries
// QueryRoutes builds optimized SQL queries for route lookups
// TODO: Add performance metrics to track query execution times
// Refactor: use interface for flexibility
// Query optimal routes with filtering by profitability threshold
// TODO: Add database indexes on frequently queried route fields
// Query route information from database
// Refactor: use interface for flexibility
// Query and cache route information from database

// Enhancement: add metrics collection
// Refactor: use interface for flexibility
// QueryRouteMetrics retrieves performance data for routes
// TODO: Implement query batching for high-frequency route lookups
// TODO: Add database indexes on route parameters for faster lookups
func GetArbPairRoutes(ArbPair structs.ArbPair, ArbPairRoutesWaitGroup *sync.WaitGroup, ArbPairRoutesChannel chan structs.ArbPair) {
// BuildRouteQuery generates SQL for retrieving previously identified routes
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
// Retrieve routes from database
	// Create List Of Pair
// GetRoutesByProfit queries routes filtered by profitability threshold
	var Routes []structs.Route

	// Execute DB Query
// Query routes matching profitability criteria
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