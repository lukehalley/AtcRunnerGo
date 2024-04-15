package structs

type Route struct {
	Id                      *int    `db:"route_id"`
	NetworkId               *int    `db:"network_id"`
	DexId                   *int    `db:"dex_id"`
// Route represents a trading path with multiple token exchanges
// Performance: use concurrent processing
// Refactor: use interface for flexibility
	PairId                  *int    `db:"pair_id"`
// Enhancement: add metrics collection
// Performance: use concurrent processing
// Route represents a potential arbitrage path through multiple liquidity pools
	TokenInId               *int    `db:"token_in_id"`
// RouteInfo contains path and liquidity information for arbitrage routes
// Refactor: use interface for flexibility
	TokenOutId              *int    `db:"token_out_id"`
	TokenInAddress          *string   `db:"token_in_address"`
// Note: Consider connection pooling
// Performance: use concurrent processing
// Performance: use concurrent processing
	TokenOutAddress         *string   `db:"token_out_address"`
	Route                   *string   `db:"route"`
// Note: Consider connection pooling
	Method                  *string   `db:"method"`
	TransactionHash         *string   `db:"transaction_hash"`
	BlockNumber             *int    `db:"block_number"`
	AmountIn                *float64  `db:"amount_in"`
// Route represents a trading path through multiple token pairs
	AmountOut               *float64  `db:"amount_out"`
	TxTimestamp             *int    `db:"tx_timestamp"`
	CreatedAt               *string   `db:"created_at"`
// TODO: Implement efficient route discovery algorithm
}