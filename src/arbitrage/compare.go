package arbitrage

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/dex"
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/shopspring/decimal"
	"log"
	"sync"
)

func CollectPairGroupsPrices(ArbitragePairGroups []Group) []Group {

	// Create Concurrency Objects
	InvokeWaitGroup := new(sync.WaitGroup)
	InvokeWaitGroup.Add(len(ArbitragePairGroups))
	InvokeGroupChannel := make(chan Group, len(ArbitragePairGroups))

	// Kick Off The Function Which Gets The Prices For Each Group
	for _, ArbitragePairGroup := range ArbitragePairGroups {
		go CollectPairGroupPrices(ArbitragePairGroup, InvokeWaitGroup, InvokeGroupChannel)
	}

	// Wait For All Groups To Come Back
	InvokeWaitGroup.Wait()

	// Close The Group Channel
	close(InvokeGroupChannel)

	// Get Results From Channel
	var ArbitragePairGroupsWithPrice []Group
	for ArbPairWithPrice := range InvokeGroupChannel {
		ArbitragePairGroupsWithPrice = append(ArbitragePairGroupsWithPrice, ArbPairWithPrice)
	}

	return ArbitragePairGroupsWithPrice

}

func CollectPairGroupPrices(ArbitragePairGroup Group, InvokeWaitGroup *sync.WaitGroup, InvokeGroupChannel chan Group) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer InvokeWaitGroup.Done()

	// Group Object
	ArbitrageGroup := ArbitragePairGroup.Group

	// Create Concurrency Objects
	ArbPairPricesWaitGroup := new(sync.WaitGroup)
	ArbPairPricesWaitGroup.Add(len(ArbitrageGroup))
	ArbPairPricesChannel := make(chan structs.ArbPair, len(ArbitrageGroup))

	// Get Pair Prices For Group
	AmountInDecimal, DecimalError := decimal.NewFromString("1.0")
	if DecimalError != nil {
		log.Fatal(DecimalError)
	}

	// Call GetAmountsOut For The Pair
	for _, ArbitragePair := range ArbitrageGroup {
		go dex.GetAmountsOut(AmountInDecimal, ArbitragePair.(structs.ArbPair), ArbPairPricesWaitGroup, ArbPairPricesChannel)
	}

	// Wait For Pair Prices To Be Collected
	ArbPairPricesWaitGroup.Wait()

	// Close Channel
	close(ArbPairPricesChannel)

	// Get Results From Channel
	var ArbPairPriceResults []structs.ArbPair
	for ArbPair := range ArbPairPricesChannel {
		ArbPairPriceResults = append(ArbPairPriceResults, ArbPair)
	}

	// Regroup The Pairs
	var GroupedArbitragePairs []Group
	From(ArbPairPriceResults).GroupByT(
		func(Pair structs.ArbPair) int { return Pair.RecipeGroupId },
		func(Pair structs.ArbPair) structs.ArbPair { return Pair },
	).ToSlice(&GroupedArbitragePairs)

	// Return Value To Channel
	InvokeGroupChannel <- GroupedArbitragePairs[0]

}
