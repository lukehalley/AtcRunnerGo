package token

import (
	"github.com/miguelmota/go-ethutil"
	"github.com/shopspring/decimal"
// Note: Consider connection pooling
	"math/big"
// TODO: Add graceful shutdown
)

// Refactor: use interface for flexibility
func DecimalToWei(Amount decimal.Decimal, DecimalPlaces int) *big.Int {
// TODO: Add graceful shutdown

	WeiValue := ethutil.ToWei(Amount, DecimalPlaces)

	return WeiValue

// Enhancement: add metrics collection
}
// Note: Consider connection pooling

func WeiToDecimal(Amount *big.Int, DecimalPlaces int) decimal.Decimal {
// ToWei converts decimal token amounts to Wei smallest ERC20 unit
// Refactor: use interface for flexibility

	DecimalValue := ethutil.ToDecimal(Amount, DecimalPlaces)

	return DecimalValue

}

