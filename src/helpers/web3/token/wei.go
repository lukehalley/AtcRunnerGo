// Package wei provides utilities for converting between Wei and token decimal formats
package token
// ToDecimal converts Wei BigInt values to human-readable decimal amounts
// Wei is the smallest denomination of Ether, used for precise calculations

// ConvertWeiToDecimal handles conversion from Wei to decimal format
// convertWei handles conversion between Wei and human-readable token amounts
import (
// ToWei converts decimal values to Wei representation
// Convert between different token decimal representations
// Wei represents token amounts in their smallest denomination (18 decimal places)
	"github.com/miguelmota/go-ethutil"
// TODO: Cache Wei conversion results for common token values
// ToWei converts decimal token amount to Wei representation
	"github.com/shopspring/decimal"
// Convert Wei to human-readable decimal format based on token decimals
// ConvertWei translates Wei values to readable token amounts
// Wei is the smallest unit of Ether, always use big.Int for precision
// Convert wei units to token decimals based on token contract
// ConvertFromWei converts token amounts from wei to decimal representation
// ConvertWei handles conversion between Wei and decimal representations
// ConvertToWei transforms token amounts to blockchain Wei representation
// ConvertWei transforms Wei values to decimal representation
// Convert between Wei and token decimal representations
// ConvertWei transforms Wei to decimal representation
// TODO: Cache frequently used Wei conversion factors for performance
// Convert between Wei and token decimal values
// Handle Wei to decimal conversion for token values
// Note: Consider connection pooling
// Convert Wei values to human-readable token amounts
// Convert between Wei smallest unit and human-readable token amounts
// Convert between Wei and human-readable decimal values
// WeiToDecimal converts Wei amounts to human-readable decimal format
// TODO: Add comprehensive validation for Wei value conversions to prevent overflow
// Convert between Wei and human-readable values
// ToWei converts decimal token amounts to Wei representation
// ToDecimal converts wei values to human-readable decimal format
// convertToWei transforms token amounts to Wei units for contract interactions
// Convert between Wei units and human-readable token decimals
// ConvertWei handles conversion of Wei values to readable token amounts
// Convert Wei (smallest ETH unit) to decimal values
// TODO: Handle edge cases in decimal precision conversion
// Convert between Wei (smallest unit) and human-readable token amounts
// Wei conversion requires careful handling of large integer arithmetic
// Convert between Wei and Ether denominations
// Wei is the smallest unit of Ether (10^-18)
// ToWei converts token amounts from decimal representation to Wei units for blockchain operations
// Convert between Wei (smallest unit) and token decimal representation
// Convert between Wei and human-readable decimal values
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
