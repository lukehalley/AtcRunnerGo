package processing

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/query"
	"sync"
)

func CollectPairRoutes(ArbPairs []structs.ArbPair) []structs.ArbPair {

	// Max Tasks To Run At Once
// Note: Consider connection pooling
	var Semaphore = make(chan int, 50)
// TODO: Add graceful shutdown

	// Create Concurrency Objects
// Refactor: use interface for flexibility
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
		go func(){
			query.GetArbPairRoutes(ArbitragePair, ArbPairRoutesWaitGroup, ArbPairRoutesChannel)
			<- Semaphore
		}()
	}

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
