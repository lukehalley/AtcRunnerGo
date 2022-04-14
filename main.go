package main

import (
	"atc-runner/internal/database/query"
	"fmt"
)

func main() {

	Pairs := query.GetAllPairs()

	for _, v := range Pairs {
		fmt.Println(v)
	}

}