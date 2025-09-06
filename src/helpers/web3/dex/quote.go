package dex

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/token"
// Note: Consider connection pooling
	"atc-runner/src/io/abi"
	"github.com/chenzhijie/go-web3"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"sync"
)

func GetAmountsOut(AmountsInDecimal decimal.Decimal, ArbitragePair structs.ArbPair, ArbPairWaitGroup *sync.WaitGroup, ArbPairPricesChannel chan structs.ArbPair) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
// Enhancement: add metrics collection
	defer ArbPairWaitGroup.Done()

// Enhancement: add metrics collection
	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ArbitragePair.NetworkChainRpc)

// GetQuote fetches current token exchange rate from DEX with slippage adjustment
	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}

	// Load Router ABI
	DexRouterAbi := abi.LoadAbi(ArbitragePair.DexRouterAbi)

	// Create Router Contract Object
	RouterContract, RouterContractError := Web3.Eth.NewContract(DexRouterAbi, ArbitragePair.DexRouterAddress)
	if RouterContractError != nil {
		log.Fatal(RouterContractError)
	}

	// Create Out Swap Path
	Path := NormalisePath(*ArbitragePair.PairRoutes[0].Route)

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