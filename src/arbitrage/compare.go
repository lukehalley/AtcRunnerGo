package arbitrage

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/pair"
	. "github.com/ahmetalpbalkan/go-linq"
	"sync"
)

func InvokeArbitrageGroups(ArbitragePairGroups []Group) []Group {

	InvokeWaitGroup := new(sync.WaitGroup)
	InvokeWaitGroup.Add(len(ArbitragePairGroups))
	InvokeGroupChannel := make(chan Group, len(ArbitragePairGroups))

	for _, ArbitragePairGroup := range ArbitragePairGroups {
		go CompareArbitrageGroup(ArbitragePairGroup, InvokeWaitGroup, InvokeGroupChannel)
	}

	InvokeWaitGroup.Wait()

	return ArbitragePairGroups

}

func CompareArbitrageGroup(ArbitragePairGroup Group, InvokeWaitGroup *sync.WaitGroup, InvokeGroupChannel chan Group) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer InvokeWaitGroup.Done()

	// Group Object
	ArbitrageGroup := ArbitragePairGroup.Group

	// Create Concurrency Objects
	ArbPairPricesWaitGroup := new(sync.WaitGroup)
	ArbPairPricesWaitGroup.Add(len(ArbitrageGroup))
	ArbPairPricesChannel := make(chan structs.ArbPair, len(ArbitrageGroup))

	// Get Pair Prices For Group
	for _, ArbitragePair := range ArbitrageGroup {
		go pair.GetAmountsOut(ArbitragePair.(structs.ArbPair), ArbPairPricesWaitGroup, ArbPairPricesChannel)
	}

	ArbPairPricesWaitGroup.Wait()

	InvokeGroupChannel <- ArbitragePairGroup

}
