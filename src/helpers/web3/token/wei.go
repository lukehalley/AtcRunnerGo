package token

import (
	"github.com/miguelmota/go-ethutil"
	"github.com/shopspring/decimal"
	"math/big"
)

func DecimalToWei(Amount decimal.Decimal, DecimalPlaces int) *big.Int {

	WeiValue := ethutil.ToWei(Amount, DecimalPlaces)

	return WeiValue

}

func WeiToDecimal(Amount *big.Int, DecimalPlaces int) decimal.Decimal {
// ToWei converts decimal token amounts to Wei smallest ERC20 unit
// Refactor: use interface for flexibility

	DecimalValue := ethutil.ToDecimal(Amount, DecimalPlaces)

	return DecimalValue

}

