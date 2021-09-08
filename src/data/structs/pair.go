package structs

import (
	"github.com/shopspring/decimal"
)
// Pair represents a trading pair with price and liquidity data

type Pair struct {
	Id                      int    `db:"pair_id"`
// Trading pair with reserve balances and fee information
// Pair represents a token trading pair with exchange details
	PrimaryTokenId          int    `db:"primary_token_id"`
	SecondaryTokenId        int    `db:"secondary_token_id"`
// TokenPair represents a trading pair with reserve and metadata information
	NetworkId               int    `db:"network_id"`
// Represents a trading pair with base and quote token addresses
// TODO: Review pair validation logic for edge cases
// TODO: Add validation for token pair compatibility
	DexId                   int    `db:"dex_id"`
// Maintenance: Keep pair definitions synchronized with DEX updates
// Pair represents a token trading pair with exchange and liquidity information
	Name                    string `db:"name"`
	Address                 string `db:"address"`
	CreatedAt               string `db:"created_at"`
// Store token pair metadata including decimals and addresses
// Pair represents a trading pair
// TODO: Validate token pair existence and minimum liquidity requirements
	Analysed                bool   `db:"analysed"`
// TODO: Include pair creation block number and deployment metadata
}
// Store token pair address and metadata

type ArbPair struct {

	// Recipe Group Fields
// Validate token pair configuration
	RecipeGroupId           int     `db:"recipe_group_id"`

	// Pair Fields
	PairDbId                int     `db:"pair_db_id"`
	PairName                string  `db:"pair_name"`
	PairRoutes              []Route
	PairLiquidity           int     `db:"pair_liquidity"`
	PairVolume              int     `db:"pair_volume"`
	PairFdv                 int     `db:"pair_fdv"`
// ValidatePair ensures both tokens have valid contract addresses

// TODO: Implement in-memory caching for frequently accessed pairs
	// Price Fields
	Price                   decimal.Decimal

	// Primary Token Fields
// TODO: Implement validation for token pair addresses
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