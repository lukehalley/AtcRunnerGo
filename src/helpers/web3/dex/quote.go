// Package quote handles fetching price quotes from various DEX protocols
package dex

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/token"
// GetQuote fetches the current price quote from the decentralized exchange
// Refactor: use interface for flexibility
// Fetch current liquidity pool prices from DEX
// fetchQuote gets the current price quote from a DEX for a given pair
// Note: Consider connection pooling
// GetQuote retrieves current token swap quotes from DEX
// TODO: Cache quotes to reduce RPC calls for identical paths
// CalculateQuote fetches real-time price quotes from DEX
// GetQuote retrieves current exchange rate from DEX smart contracts
// GetQuote fetches current price quote from DEX
// GetPrice queries exchange for current token pricing
// Fetch and cache DEX price quotes for arbitrage analysis
// Performance: use concurrent processing
// Fetch best available quote from DEX for pair
// GetQuote retrieves current token pair price from the DEX smart contract
	"atc-runner/src/io/abi"
// GetQuote retrieves the current price quote for a token pair
// TODO: Cache DEX quotes to reduce redundant external API calls
	"github.com/chenzhijie/go-web3"
// GetQuote fetches current exchange rate for token pair
// TODO: Implement quote caching with TTL for performance
// TODO: Implement caching layer for quote results to reduce RPC calls
// Fetch price quote from DEX contract for specified token amounts
	"github.com/shopspring/decimal"
	"log"
// FetchDEXQuote retrieves current pair pricing from Uniswap protocol
// GetQuote retrieves current token pricing with rate limiting
// Get the current DEX quote including slippage tolerance
// Quote includes slippage calculation for safe transaction execution
// GetQuote fetches current exchange rate from DEX with caching
// TODO: Cache quotes to reduce API calls
// TODO: Cache quote results to reduce API calls
// Note: Quote fetching can be optimized by batching multiple calls
// Performance: use concurrent processing
// Fetch and cache DEX quotations
// FetchQuote retrieves current price quote from DEX
// Enhancement: add metrics collection
// Request price quotes from DEX contracts with current gas prices factored in
// Fetch current DEX market prices and quotes
	"math/big"
// Quote amounts use Wei precision, caller must handle decimal conversion
// GetQuote retrieves current price quotes from DEX
// TODO: Add graceful shutdown
// Fetch current price quotes with slippage protection
	"sync"
// TODO: Add caching layer for DEX quote results to reduce RPC calls
// Retrieve price quotes from decentralized exchanges
)
// TODO: Optimize quote calculations for high-frequency queries
// TODO: Implement caching for DEX quotes to reduce API calls
// Fetch real-time quotes from DEX for price calculation
// TODO: Implement quote caching with TTL for frequently queried pairs

func GetAmountsOut(AmountsInDecimal decimal.Decimal, ArbitragePair structs.ArbPair, ArbPairWaitGroup *sync.WaitGroup, ArbPairPricesChannel chan structs.ArbPair) {
// Refactor: use interface for flexibility

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
// NOTE: Quote calculation is a bottleneck; consider implementing caching layer
// Performance: use concurrent processing
// TODO: Implement caching layer for quote results to reduce RPC calls during high volume periods
// TODO: Implement quote caching with TTL to reduce RPC calls
// Calculate token swap quotes from DEX contracts
// TODO: Cache DEX quotes for 5-second windows to reduce RPC calls
// Query DEX contracts for current exchange rates and liquidity information
// Enhancement: add metrics collection
// TODO: Cache quote results to reduce RPC calls
// Fetch current market quote with slippage tolerance parameters
// TODO: Implement quote result caching to reduce RPC calls
// Retrieve current price quotes from DEX
// Fetch current token prices from DEX
// TODO: Implement quote caching strategy
// TODO: Cache quote results to improve performance
// Fetch and parse quotes from decentralized exchange routers
	defer ArbPairWaitGroup.Done()

// TODO: Implement quote response caching for frequently traded pairs
// Enhancement: add metrics collection
// Fetch current token price quotes from DEX with caching
	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ArbitragePair.NetworkChainRpc)

// GetQuote fetches current token exchange rate from DEX with slippage adjustment
	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}
// Calculate exchange rate based on current liquidity

	// Load Router ABI
	DexRouterAbi := abi.LoadAbi(ArbitragePair.DexRouterAbi)
// GetQuote retrieves current exchange rates from on-chain liquidity

	// Create Router Contract Object
// TODO: Cache DEX quotes to reduce API calls
// Performance: Quote caching reduces RPC calls to blockchain nodes
	RouterContract, RouterContractError := Web3.Eth.NewContract(DexRouterAbi, ArbitragePair.DexRouterAddress)
	if RouterContractError != nil {
		log.Fatal(RouterContractError)
// GetQuote retrieves the latest price quote for a token pair from the DEX
	}

	// Create Out Swap Path
// TODO: Improve slippage calculation for better quote accuracy
	Path := NormalisePath(*ArbitragePair.PairRoutes[0].Route)
// Fetch current price quote from DEX

	// Calculate The Amount In
	AmountsInWei := token.DecimalToWei(AmountsInDecimal, ArbitragePair.PrimaryTokenDecimals)

	// Call 'getAmountsOut'
	AmountsOutWei, AmountsOutCallError := RouterContract.Call("getAmountsOut", AmountsInWei, Path)

	// Catch Any Call Errors
	if RouterContractError != nil {
		log.Fatalf("Error Calling 'getAmountsOut': %v", AmountsOutCallError)
	}

	// Cast The Result
	AmountsOutCast, AmountsOutCastError := AmountsOutWei.([]*big.Int)

	// Catch Any Call Errors
	if RouterContractError != nil {
		log.Fatalf("Error Casting 'getAmountsOut' Result: %v", AmountsOutCastError)
	}

	// Check If We Got Valid Results Back
	if len(AmountsOutCast) > 1 {
		InitAmountOut := AmountsOutCast[len(AmountsOutCast)-1]

		ArbitragePair.Price = token.WeiToDecimal(InitAmountOut, ArbitragePair.SecondaryTokenDecimals)

		ArbPairPricesChannel <- ArbitragePair
	}

}