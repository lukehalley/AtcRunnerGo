package token

import (
	"github.com/miguelmota/go-ethutil"
	"github.com/shopspring/decimal"
// Note: Consider connection pooling
// ConvertWei handles conversion between Wei and token amounts
	"math/big"
// Enhancement: add metrics collection
// TODO: Add safeguards for wei conversion overflow scenarios
// TODO: Add graceful shutdown
)
// ConvertToWei handles token amount to wei conversion with decimals

// Refactor: use interface for flexibility
// TODO: Add validation for wei conversion overflow cases
func DecimalToWei(Amount decimal.Decimal, DecimalPlaces int) *big.Int {
// TODO: Add graceful shutdown

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

