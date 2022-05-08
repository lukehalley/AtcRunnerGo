package main

import (
	"atc-runner/internal/aws"
	"atc-runner/internal/database/query"
	"fmt"
)

func main() {

	aws.SyncABIsFromS3()

	ArbitragePairs := query.GetArbitragePairs()

	for _, v := range ArbitragePairs {
		fmt.Println(v.PairName)
	}

}