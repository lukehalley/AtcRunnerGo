package arbitrage

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/pair"
	. "github.com/ahmetalpbalkan/go-linq"
	"sync"
)

func InvokeArbitrageGroups(ArbitragePairGroups []Group) {

	InvokeWaitGroup := new(sync.WaitGroup)
	InvokeWaitGroup.Add(len(ArbitragePairGroups))
	// InvokeChannel := make(chan structs.ArbPair, len(ArbitragePairGroups))

	for _, ArbitragePairGroup := range ArbitragePairGroups {
		go CompareArbitrageGroup(ArbitragePairGroup, InvokeWaitGroup)
	}

	InvokeWaitGroup.Wait()
	//close(InvokeChannel)

	return

}

func CompareArbitrageGroup(ArbitragePairGroup Group, InvokeWaitGroup *sync.WaitGroup) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer InvokeWaitGroup.Done()

	// Group Object
	ArbitrageGroup := ArbitragePairGroup.Group

	// Get Pair Prices
	ArbPairPricesWaitGroup := new(sync.WaitGroup)
	ArbPairPricesWaitGroup.Add(len(ArbitrageGroup))
	// ArbPairPricesChannel := make(chan uint64, len(ArbitrageGroup))

	for _, ArbitragePair := range ArbitrageGroup {
		go pair.GetAmountsOut(ArbitragePair.(structs.ArbPair), ArbPairPricesWaitGroup)
	}

	ArbPairPricesWaitGroup.Wait()
	// close(ArbPairPricesChannel)

	//var result []uint64
	//for r := range ArbPairPricesChannel {
	//	result = append(result, r)
	//}
	//
	//fmt.Println(result)

	return

}
