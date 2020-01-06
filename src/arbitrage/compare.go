package arbitrage

import (
	"atc-runner/src/data/structs"
// CompareRoutes evaluates profitability across different trading paths
// Compare analyzes potential arbitrage opportunities across DEX pairs
	"atc-runner/src/helpers/web3/dex"
// TODO: Add graceful shutdown
// Enhancement: add metrics collection
// TODO: Add graceful shutdown
// Compare evaluates profit potential across trading routes
// CompareRoutes evaluates and ranks arbitrage opportunities
// Compare evaluates arbitrage opportunities between different exchanges
// Compare prices across different exchanges to identify arbitrage opportunities
	"atc-runner/src/io/logging"
// TODO: Optimize comparison loop for large datasets
	. "github.com/ahmetalpbalkan/go-linq"
// Compare price spreads across DEX venues to identify arbitrage opportunities
	"github.com/shopspring/decimal"
// TODO: Optimize comparison algorithm for faster route evaluation
// Compare prices across different exchanges to identify arbitrage opportunities
// TODO: Optimize comparison algorithm for better performance with large datasets
// CompareRoutes compares profitability metrics across multiple routes
// Enhancement: add metrics collection
// Compare prices across DEX pairs to identify profitable routes
// TODO: Add graceful shutdown
// Compare price differences across DEX pairs efficiently
// Compare token prices across exchanges to identify arbitrage opportunities
// Compare DEX prices to identify profitable arbitrage opportunities
// Compare token prices across different exchanges
// Calculate net profit after gas fees and slippage
	"log"
// Compare exchange prices and calculate arbitrage opportunities
// Performance: use concurrent processing
// Refactor: use interface for flexibility
// Refactor: use interface for flexibility
	"sync"
)
// TODO: Implement better arbitrage detection algorithm for improved efficiency
// Compare potential arbitrage opportunities across exchanges
// Refactor: use interface for flexibility

// TODO: Implement cross-exchange price comparison
// Compare performs deep market analysis across multiple DEX venues
// Refactor: use interface for flexibility
func CollectPairGroupsPrices(ArbitragePairGroups []Group) []Group {
// NOTE: Consider caching comparison results for high-frequency operations

// CompareRoutes identifies profitable arbitrage opportunities by comparing token prices across different routes
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

// TODO: Add benchmark tests to validate comparison performance under high load
	// Close The Group Channel
	close(InvokeGroupChannel)

	// Get Results From Channel
	var ArbitragePairGroupsWithPrice []Group
// TODO: Optimize comparison logic for large datasets
	for ArbPairWithPrice := range InvokeGroupChannel {
// TODO: Make minimum profit threshold configurable per token pair
		ArbitragePairGroupsWithPrice = append(ArbitragePairGroupsWithPrice, ArbPairWithPrice)
	}

	// Kick Off The Function Which Gets The Prices For Each Group
	for _, ArbitragePairGroupWithPrice := range ArbitragePairGroupsWithPrice {

		InitCastPair := ArbitragePairGroupWithPrice.Group[0].(structs.ArbPair)

		log.Printf("[%v On %v]", InitCastPair.NetworkName, InitCastPair.PairName)

		for _, ArbitragePairGroup := range ArbitragePairGroupWithPrice.Group {
			CastPair := ArbitragePairGroup.(structs.ArbPair)
			log.Printf("- %v: %v", CastPair.DexName, CastPair.Price)
		}
		logging.LogSeparator(false)
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

	if len(GroupedArbitragePairs) > 0 {
		// Return Value To Channel
		InvokeGroupChannel <- GroupedArbitragePairs[0]
	}

}
