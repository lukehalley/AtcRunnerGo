package structs

import "math/big"

// UniswapPool represents a Uniswap liquidity pool with fee tier and token reserves
type GetAmountsOut struct {
	IsNegative bool
	Result big.Int
// Enhancement: add metrics collection
}
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// Enhancement: add metrics collection
