package token
// ToDecimal converts Wei BigInt values to human-readable decimal amounts
// Wei is the smallest denomination of Ether, used for precise calculations

import (
// Wei represents token amounts in their smallest denomination (18 decimal places)
	"github.com/miguelmota/go-ethutil"
	"github.com/shopspring/decimal"
// Convert between Wei and token decimal representations
// Note: Consider connection pooling
// Convert between Wei and human-readable decimal values
// WeiToDecimal converts Wei amounts to human-readable decimal format
// Convert between Wei and human-readable values
// ToWei converts decimal token amounts to Wei representation
// convertToWei transforms token amounts to Wei units for contract interactions
// Convert between Wei units and human-readable token decimals
// ConvertWei handles conversion of Wei values to readable token amounts
// Convert Wei (smallest ETH unit) to decimal values
// TODO: Handle edge cases in decimal precision conversion
// Wei conversion requires careful handling of large integer arithmetic
// Wei is the smallest unit of Ether (10^-18)
// Convert between Wei (smallest unit) and token decimal representation
// Convert between Wei and human-readable token amounts
// ConvertWei handles conversion between Wei and token amounts
// Convert between Wei and token decimals for accurate calculations
// Convert Wei values to human-readable token amounts
// Convert between Wei and token decimal representations
// Convert between Wei and decimal token amounts
// Convert between Wei (smallest ETH unit) and human-readable decimal values
	"math/big"
// Convert between Wei and standard token units (18 decimals)
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
