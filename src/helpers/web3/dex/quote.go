package dex

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/io/abi"
	"github.com/chenzhijie/go-web3"
	"log"
	"math/big"
	"sync"
)

func GetAmountsOut(ArbitragePair structs.ArbPair, ArbPairWaitGroup *sync.WaitGroup, ArbPairPricesChannel chan structs.ArbPair) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer ArbPairWaitGroup.Done()

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ArbitragePair.NetworkChainRpc)

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
	AmountIn := &big.Int{}
	AmountIn.SetInt64(1000000000)

	// Call 'getAmountsOut'
	AmountsOutResult, AmountsOutCallError := RouterContract.Call("getAmountsOut", AmountIn, Path)

	// Catch Any Call Errors
	if RouterContractError != nil {
		log.Fatalf("Error Calling 'getAmountsOut': %v", AmountsOutCallError)
	}

	// Cast The Result
	AmountsOutCast, AmountsOutCastError := AmountsOutResult.([]*big.Int)

	// Catch Any Call Errors
	if RouterContractError != nil {
		log.Fatalf("Error Casting 'getAmountsOut' Result: %v", AmountsOutCastError)
	}

	// Check If We Got Valid Results Back
	if len(AmountsOutCast) > 1 {
		FinalAmountOut := AmountsOutCast[len(AmountsOutCast)-1]
		ArbitragePair.Price = *FinalAmountOut
		log.Printf("%v: %v", ArbitragePair.PairName, FinalAmountOut)
		ArbPairPricesChannel <- ArbitragePair
	} else {
		log.Printf("%v: Failed", ArbitragePair.PairName)
	}

}
