package structs

import "math/big"

type GetAmountsOut struct {
	IsNegative bool
	Result big.Int
}
