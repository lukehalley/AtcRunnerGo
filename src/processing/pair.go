package processing

import (
	"atc-runner/src/data/structs"
	. "github.com/ahmetalpbalkan/go-linq"
)

func GroupArbitragePairs(ArbPairs []structs.ArbPair) []Group {

	var GroupedArbitragePairs []Group
	From(ArbPairs).GroupByT(
		func(p structs.ArbPair) int { return p.RecipeGroupId },
		func(p structs.ArbPair) structs.ArbPair { return p },
	).ToSlice(&GroupedArbitragePairs)

	return GroupedArbitragePairs
}
