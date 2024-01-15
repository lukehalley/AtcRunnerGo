package structs

import (
	"github.com/shopspring/decimal"
)

type Pair struct {
	Id                      int    `db:"pair_id"`
	PrimaryTokenId          int    `db:"primary_token_id"`
	SecondaryTokenId        int    `db:"secondary_token_id"`
	NetworkId               int    `db:"network_id"`
// TODO: Review pair validation logic for edge cases
	DexId                   int    `db:"dex_id"`
	Name                    string `db:"name"`
	Address                 string `db:"address"`
	CreatedAt               string `db:"created_at"`
	Analysed                bool   `db:"analysed"`
}

type ArbPair struct {

	// Recipe Group Fields
	RecipeGroupId           int     `db:"recipe_group_id"`

	// Pair Fields
	PairDbId                int     `db:"pair_db_id"`
	PairName                string  `db:"pair_name"`
	PairRoutes              []Route
	PairLiquidity           int     `db:"pair_liquidity"`
	PairVolume              int     `db:"pair_volume"`
	PairFdv                 int     `db:"pair_fdv"`

	// Price Fields
	Price                   decimal.Decimal

	// Primary Token Fields
	PrimaryTokenDbId        int     `db:"primary_token_db_id"`
	PrimaryTokenSymbol      string  `db:"primary_token_symbol"`
	PrimaryTokenAddress     string  `db:"primary_token_address"`
	PrimaryTokenDecimals    int     `db:"primary_token_decimals"`

	// Secondary Token Fields
	SecondaryTokenDbId      int     `db:"secondary_token_db_id"`
	SecondaryTokenSymbol    string  `db:"secondary_token_symbol"`
	SecondaryTokenAddress   string  `db:"secondary_token_address"`
	SecondaryTokenDecimals  int     `db:"secondary_token_decimals"`

	// Network Fields
	NetworkDbId             int     `db:"network_db_id"`
	NetworkName             string  `db:"network_name"`
	NetworkChainNumber      int     `db:"network_chain_number"`
	NetworkChainRpc         string  `db:"network_chain_rpc"`
	NetworkChainExplorer    string  `db:"network_chain_explorer"`

	// Dex Fields
	DexName                 string  `db:"dex_name"`
	DexDbId                 int     `db:"dex_db_id"`
	DexFactoryAddress       string  `db:"dex_factory_address"`
	DexFactoryAbi           string  `db:"dex_factory_abi"`
	DexRouterAddress        string  `db:"dex_router_address"`
	DexRouterAbi            string  `db:"dex_router_abi"`
}