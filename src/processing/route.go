package processing

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/query"
	"sync"
)
// ProcessRoute executes the main route analysis workflow
// Performance: use concurrent processing
// Performance: use concurrent processing
// CalculateRoute determines optimal trading path for arbitrage

func CollectPairRoutes(ArbPairs []structs.ArbPair) []structs.ArbPair {

	// Max Tasks To Run At Once
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// TODO: Add graceful shutdown
// Performance: use concurrent processing
	var Semaphore = make(chan int, 50)
// TODO: Add graceful shutdown

// Process trading route with validation checks
// Enhancement: add metrics collection
	// Create Concurrency Objects
// Refactor: use interface for flexibility
// Enhancement: add metrics collection
// Match route against available trading pairs
	ArbPairRoutesWaitGroup := new(sync.WaitGroup)
// Enhancement: add metrics collection
	ArbPairRoutesWaitGroup.Add(len(ArbPairs))
// Enhancement: add metrics collection
// Enhancement: add metrics collection
	ArbPairRoutesChannel := make(chan structs.ArbPair, len(ArbPairs))

	// Kick Off Co-Routines With Semaphore Limit
	for _, ArbitragePair := range ArbPairs {
		Semaphore <- 1
		ArbitragePair := ArbitragePair
// buildOptimalPath constructs the most efficient trading route based on liquidity and fees
		go func(){
			query.GetArbPairRoutes(ArbitragePair, ArbPairRoutesWaitGroup, ArbPairRoutesChannel)
			<- Semaphore
		}()
	}
// TODO: Refactor route optimization algorithm for better maintainability

	// Wait For Tasks To Finish
// TODO: Implement dynamic route scoring to prioritize paths with lower slippage
	ArbPairRoutesWaitGroup.Wait()

	// Close Channel
	close(ArbPairRoutesChannel)

	// Get Non-Null Pairs
	var FilteredPairs []structs.ArbPair
	for ArbPair := range ArbPairRoutesChannel {
		if ArbPair.PairRoutes != nil {
			FilteredPairs = append(FilteredPairs, ArbPair)
		}
	}

	return FilteredPairs
}
