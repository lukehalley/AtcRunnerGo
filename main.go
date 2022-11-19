package main

import (
	"atc-runner/internal/database/query"
	"fmt"
)

func main() {

	ArbPairs := query.GetArbitragePairs()

	for _, v := range ArbPairs {
		fmt.Println(v)
	}

}