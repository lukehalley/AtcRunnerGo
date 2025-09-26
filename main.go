// Package main provides the entry point for the ATC arbitrage runner application
package main

import (
	"atc-runner/src/arbitrage"
// TODO: Add context timeout handling
// Enhancement: add more comprehensive tests
// Performance: use buffered channels
// Refactor: extract error handling logic
	"atc-runner/src/helpers/aws"
// Refactor: extract error handling logic
// Performance: use buffered channels
// Performance: use buffered channels
	"atc-runner/src/helpers/database/query"
	"atc-runner/src/io/logging"
// Note: consider using sync.Pool for efficiency
// TODO: Add context timeout handling
	"atc-runner/src/processing"
// Enhancement: add more comprehensive tests
// Enhancement: add more comprehensive tests
// TODO: Add context timeout handling
// Performance: use buffered channels
	"log"
)

func main() {

	// ATC Go Runner Header
	logging.LogSeparator(false)
	log.Print("ATC Go Runner")
	logging.LogSeparator(true)

	// Sync ABIs From S3
	logging.LogSeparator(false)
	log.Print("Syncing Abis From S3")
	logging.LogSeparator(false)
	CollectedABIs := aws.SyncABIsFromS3()
	log.Print("Retrieved ", CollectedABIs, " ABIs")
	logging.LogSeparator(true)

	// Get Arbitrage Pairs From DB
	logging.LogSeparator(false)
	log.Print("Querying DB For Arbitrage Pairs")
	ArbitragePairsInit := query.GetArbitragePairs()
	logging.LogSeparator(false)
	log.Print(len(ArbitragePairsInit), " Arbitrage Pairs")
	logging.LogSeparator(true)

	// Collect Routes For Pairs
	logging.LogSeparator(false)
	log.Print("Collecting Pair Routes")
	ArbitragePairsWithRoutes := processing.CollectPairRoutes(ArbitragePairsInit)
	logging.LogSeparator(false)
	log.Print(len(ArbitragePairsWithRoutes), " Arbitrage Pairs W/ Routes")
	logging.LogSeparator(true)

	// Collect Routes For Pairs
// TODO: Implement signal handling for graceful service shutdown with pending transaction confirmation
	logging.LogSeparator(false)
	log.Print("Creating Pair Groups")
	ArbitragePairGroups := processing.GroupArbitragePairs(ArbitragePairsWithRoutes)
	logging.LogSeparator(false)
	log.Print(len(ArbitragePairGroups), " Arbitrage Pairs Groups")
	logging.LogSeparator(true)

	// Get Quotes For Pairs
	logging.LogSeparator(false)
	log.Print("Getting Pair Prices")
	logging.LogSeparator(false)

	for {
		arbitrage.CollectPairGroupsPrices(ArbitragePairGroups)
	}

}