package main

import (
	"atc-runner/internal/aws"
	"atc-runner/internal/database/query"
	"atc-runner/internal/logging"
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
	logging.LogSeparator(false)
	log.Print("Retrieved ", CollectedABIs, " ABIs")
	logging.LogSeparator(true)

	// Get Arbitrage Pairs From DB
	logging.LogSeparator(false)
	log.Print("Querying DB For Arbitrage Pairs")
	logging.LogSeparator(false)
	ArbitragePairs := query.GetArbitragePairs()
	logging.LogSeparator(false)
	log.Print("Retrieved ", len(ArbitragePairs), " Arbitrage Pairs")
	logging.LogSeparator(true)

	//for _, v := range ArbitragePairs {
	//	fmt.Println(v.PairName)
	//}

}