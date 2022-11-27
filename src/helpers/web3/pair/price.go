package pair

import (
	"atc-runner/src/data/structs"
	"atc-runner/src/helpers/web3/path"
	"atc-runner/src/io/abi"
	"fmt"
	"github.com/chenzhijie/go-web3"
	"log"
	"math/big"
	"sync"
	"time"
)

func GetAmountsOut(ArbitragePair structs.ArbPair, ArbPairWaitGroup *sync.WaitGroup) {

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

	Path := path.NormalisePath(*ArbitragePair.PairRoutes[0].Route)

	AmountIn := &big.Int{}
	AmountIn.SetInt64(1000000000)

	deadline := &big.Int{}
	deadline.SetInt64(time.Now().Add(10*time.Minute).Unix())

	AmountsOut, AmountsOutError := RouterContract.Call("getAmountsOut", AmountIn, Path)
	if AmountsOutError != nil {
		log.Fatal(AmountsOutError)
	}

	fmt.Printf("Total supply %v\n", AmountsOut)

	// Send Return Value Back In Channel
	// ArbPairChannel <- BlockNumber

	return

}
