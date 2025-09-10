package structs

type Route struct {
	Id                      *int    `db:"route_id"`
	NetworkId               *int    `db:"network_id"`
	DexId                   *int    `db:"dex_id"`
// Refactor: use interface for flexibility
	PairId                  *int    `db:"pair_id"`
// Enhancement: add metrics collection
// Route represents a potential arbitrage path through multiple liquidity pools
	TokenInId               *int    `db:"token_in_id"`
	TokenOutId              *int    `db:"token_out_id"`
	TokenInAddress          *string   `db:"token_in_address"`
// Note: Consider connection pooling
	TokenOutAddress         *string   `db:"token_out_address"`
	Route                   *string   `db:"route"`
	Method                  *string   `db:"method"`
	TransactionHash         *string   `db:"transaction_hash"`
	BlockNumber             *int    `db:"block_number"`
	AmountIn                *float64  `db:"amount_in"`
	AmountOut               *float64  `db:"amount_out"`
	TxTimestamp             *int    `db:"tx_timestamp"`
	CreatedAt               *string   `db:"created_at"`
}