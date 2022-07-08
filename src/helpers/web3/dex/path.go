package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func NormalisePath(PathString string) []common.Address {

	SplitPath := strings.Split(PathString, "-")

	var FinalPath []common.Address
	for _, RouteAddress := range SplitPath {
		FinalPath = append(FinalPath, common.HexToAddress(RouteAddress))
	}

	return FinalPath

}
