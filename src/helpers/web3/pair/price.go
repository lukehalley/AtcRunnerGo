package pair

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/io/abi"
	"fmt"
	"github.com/chenzhijie/go-web3"
	"log"
	"sync"
)

func GetAmountsOut(ArbitragePair structs.ArbPair, ArbPairWaitGroup *sync.WaitGroup, ArbPairChannel chan uint64) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer ArbPairWaitGroup.Done()

	// Create Instance Of Web3
	Web3, Web3Error := web3.NewWeb3(ArbitragePair.NetworkChainRpc)

	// Catch Creating Web3Object
	if Web3Error != nil {
		log.Fatal(Web3Error)
	}

	// Get Block Number
	BlockNumber, BlockNumberError := Web3.Eth.GetBlockNumber()

	// Catch Getting Block Number
	if BlockNumberError != nil {
		log.Fatal(BlockNumberError)
	}

	// Load Router ABI
	DexRouterAbi := abi.LoadAbi(ArbitragePair.DexRouterAbi)

	// Create Router Contract Object
	RouterContract, RouterContractError := Web3.Eth.NewContract(DexRouterAbi, ArbitragePair.DexRouterAddress)
	if RouterContractError != nil {
		log.Fatal(RouterContractError)
	}

	AmountsOut, AmountsOutError := RouterContract.Call("getAmountsOut", 1000000000)
	if AmountsOutError != nil {
		log.Fatal(AmountsOutError)
	}

	fmt.Printf("Total supply %v\n", AmountsOut)

	// Send Return Value Back In Channel
	ArbPairChannel <- BlockNumber

}
