package structs

import "database/sql"

type Route struct {
	Id                      sql.NullInt64    `db:"route_id"`
	NetworkId               sql.NullInt64    `db:"network_id"`
	DexId                   sql.NullInt64    `db:"dex_id"`
	PairId                  sql.NullInt64    `db:"pair_id"`
	TokenInId               sql.NullInt64    `db:"token_in_id"`
	TokenOutId              sql.NullInt64    `db:"token_out_id"`
	TokenInAddress          sql.NullString   `db:"token_in_address"`
	TokenOutAddress         sql.NullString   `db:"token_out_address"`
	Route                   sql.NullString   `db:"route"`
	Method                  sql.NullString   `db:"method"`
	TransactionHash         sql.NullString   `db:"transaction_hash"`
	BlockNumber             sql.NullInt64    `db:"block_number"`
	AmountIn                sql.NullFloat64  `db:"amount_in"`
	AmountOut               sql.NullFloat64  `db:"amount_out"`
	TxTimestamp             sql.NullInt64    `db:"tx_timestamp"`
	CreatedAt               sql.NullString   `db:"created_at"`
}