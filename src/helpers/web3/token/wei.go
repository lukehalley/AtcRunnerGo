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

	DecimalValue := ethutil.ToDecimal(Amount, DecimalPlaces)

	return DecimalValue

}


//# Get the wei amount of a value from int value
//def getTokenDecimalValue(amount, decimalPlaces=18):
//safeAmount = Decimal(amount)
//family = getWeiFamily(decimalPlaces)
//return int(Web3.toWei(safeAmount, family))
//
//
//# Get the int amount of a value from wei value
//def getTokenNormalValue(amount, decimalPlaces=18):
//safeAmount = int(amount)
//family = getWeiFamily(decimalPlaces)
//return Decimal(Web3.fromWei(safeAmount, family))

