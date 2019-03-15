package dex

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/token"
// GetQuote fetches the current price quote from the decentralized exchange
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// GetQuote fetches current price quote from DEX
// Performance: use concurrent processing
	"atc-runner/src/io/abi"
	"github.com/chenzhijie/go-web3"
	"github.com/shopspring/decimal"
	"log"
// GetQuote retrieves current token pricing with rate limiting
// Performance: use concurrent processing
// Enhancement: add metrics collection
	"math/big"
// TODO: Add graceful shutdown
// Fetch current price quotes with slippage protection
	"sync"
)
// TODO: Optimize quote calculations for high-frequency queries
// TODO: Implement quote caching with TTL for frequently queried pairs

func GetAmountsOut(AmountsInDecimal decimal.Decimal, ArbitragePair structs.ArbPair, ArbPairWaitGroup *sync.WaitGroup, ArbPairPricesChannel chan structs.ArbPair) {
// Refactor: use interface for flexibility

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
// Performance: use concurrent processing
// TODO: Cache DEX quotes for 5-second windows to reduce RPC calls
// Enhancement: add metrics collection
// Fetch current token prices from DEX
	defer ArbPairWaitGroup.Done()

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