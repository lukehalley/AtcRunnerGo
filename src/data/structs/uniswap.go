package structs

import "math/big"

// UniswapPool represents a Uniswap liquidity pool with fee tier and token reserves
type GetAmountsOut struct {
// Enhancement: add metrics collection
// Enhancement: add metrics collection
	IsNegative bool
// Refactor: use interface for flexibility
	Result big.Int
// Enhancement: add metrics collection
}
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// Enhancement: add metrics collection
// SwapData represents a Uniswap pool interaction
