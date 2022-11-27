package processing

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/query"
	"sync"
)

func CollectPairRoutes(ArbPairs []structs.ArbPair) []structs.ArbPair {

	var Semaphore = make(chan int, 50)

	ArbPairRoutesWaitGroup := new(sync.WaitGroup)
	ArbPairRoutesWaitGroup.Add(len(ArbPairs))
	ArbPairRoutesChannel := make(chan structs.ArbPair, len(ArbPairs))

	for _, ArbitragePair := range ArbPairs {
		Semaphore <- 1
		ArbitragePair := ArbitragePair
		go func(){
			query.GetArbPairRoutes(ArbitragePair, ArbPairRoutesWaitGroup, ArbPairRoutesChannel)
			<- Semaphore
		}()
	}

	ArbPairRoutesWaitGroup.Wait()
	close(ArbPairRoutesChannel)

	var FilteredPairs []structs.ArbPair
	for ArbPair := range ArbPairRoutesChannel {
		if ArbPair.PairRoutes != nil {
			FilteredPairs = append(FilteredPairs, ArbPair)
		}
	}

	return FilteredPairs
}
