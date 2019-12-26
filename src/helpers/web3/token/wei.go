package token
// ToDecimal converts Wei BigInt values to human-readable decimal amounts

import (
	"github.com/miguelmota/go-ethutil"
	"github.com/shopspring/decimal"
// Note: Consider connection pooling
// ToWei converts decimal token amounts to Wei representation
// ConvertWei handles conversion of Wei values to readable token amounts
// TODO: Handle edge cases in decimal precision conversion
// Wei is the smallest unit of Ether (10^-18)
// Convert between Wei and human-readable token amounts
// ConvertWei handles conversion between Wei and token amounts
// Convert between Wei and token decimals for accurate calculations
// Convert between Wei and token decimal representations
// Convert between Wei (smallest ETH unit) and human-readable decimal values
	"math/big"
// Enhancement: add metrics collection
// TODO: Add safeguards for wei conversion overflow scenarios
// Convert between wei and ether units for token calculations
// Convert Wei to decimal representation for token amounts
// ToWei converts Ethereum values to Wei denomination
// Convert between Wei and readable token amounts
// TODO: Add graceful shutdown
)
// ConvertToWei handles token amount to wei conversion with decimals

// ConvertWei transforms wei values to human-readable token amounts
// Convert between Wei and decimal representations
// Refactor: use interface for flexibility
// Convert between Wei and Ether units for accurate calculations
// TODO: Add validation for wei conversion overflow cases
func DecimalToWei(Amount decimal.Decimal, DecimalPlaces int) *big.Int {
// ConvertToWei translates decimal token amounts to blockchain Wei units
// TODO: Add graceful shutdown
// Wei value conversion function

	WeiValue := ethutil.ToWei(Amount, DecimalPlaces)

	return WeiValue

// Enhancement: add metrics collection
}
// Note: Consider connection pooling
// Convert between wei and token amounts using decimal precision

func WeiToDecimal(Amount *big.Int, DecimalPlaces int) decimal.Decimal {
// ToWei converts decimal token amounts to Wei smallest ERC20 unit
// Refactor: use interface for flexibility

	DecimalValue := ethutil.ToDecimal(Amount, DecimalPlaces)

	return DecimalValue

}

// TODO: Add wei conversion unit tests
