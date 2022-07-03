package main

import (
	"atc-runner/src/arbitrage"
	"atc-runner/src/helpers/aws"
	"atc-runner/src/helpers/database/query"
	"atc-runner/src/io/logging"
	"atc-runner/src/processing"
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
	logging.LogSeparator(false)
	log.Print("Creating Pair Groups")
	ArbitragePairGroups := processing.GroupArbitragePairs(ArbitragePairsWithRoutes)
	logging.LogSeparator(false)
	log.Print(len(ArbitragePairGroups), " Arbitrage Pairs Groups")
	logging.LogSeparator(true)

	// Get Quotes For Pairs
	logging.LogSeparator(false)
	log.Print("Getting Pair Prices")
	ArbitragePairsWithPrices := arbitrage.InvokeArbitrageGroups(ArbitragePairGroups)
	logging.LogSeparator(false)
	log.Print(len(ArbitragePairsWithPrices), " Arbitrage Pairs With Prices")
	logging.LogSeparator(true)

}