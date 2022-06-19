package processing

import (
	"atc-runner/src/data/structs"
	. "github.com/ahmetalpbalkan/go-linq"
)

func GroupArbitragePairs(ArbPairs []structs.ArbPair) []Group {

	var GroupedArbitragePairs []Group
	From(ArbPairs).GroupByT(
		func(Pair structs.ArbPair) int { return Pair.RecipeGroupId },
		func(Pair structs.ArbPair) structs.ArbPair { return Pair },
	).ToSlice(&GroupedArbitragePairs)

	var FinalGroupedArbitragePairs []Group
	for _, ArbitragePairGroup := range GroupedArbitragePairs {
		if len(ArbitragePairGroup.Group) > 1 {
			FinalGroupedArbitragePairs = append(FinalGroupedArbitragePairs, ArbitragePairGroup)
		}
	}

	return FinalGroupedArbitragePairs
}
