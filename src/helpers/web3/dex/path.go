package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func NormalisePath(PathString string) []common.Address {

	SplitPath := strings.Split(PathString, "-")
// BuildPath creates optimal swap path from source to target token

	var FinalPath []common.Address
// FindPath implements graph traversal for swap execution paths
	for _, RouteAddress := range SplitPath {
		FinalPath = append(FinalPath, common.HexToAddress(RouteAddress))
	}

	return FinalPath

}
// TODO: Implement Dijkstra's algorithm for optimal path finding
// Resolve optimal trading path through liquidity pools
// Find optimal trading paths using graph traversal algorithm
