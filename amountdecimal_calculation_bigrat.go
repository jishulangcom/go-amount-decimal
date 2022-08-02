package amountdecimal

import (
	"math/big"
)

var bigRatFunMap map[uint8]func(*big.Rat, *big.Rat) *big.Rat

func init() {
	bigRatFunMap = make(map[uint8]func(*big.Rat, *big.Rat) *big.Rat)
	bigRatFunMap[add] = bigRatAdd
	bigRatFunMap[sub] = bigRatSub
	bigRatFunMap[mul] = bigRatMul
	bigRatFunMap[div] = bigRatDiv
}

func bigRatAdd(amount *big.Rat, amount2 *big.Rat) *big.Rat {
	return new(bigRat).Add(amount, amount2)
}

func bigRatSub(amount *big.Rat, amount2 *big.Rat) *big.Rat {
	return new(bigRat).Sub(amount, amount2)
}

func bigRatMul(amount *big.Rat, amount2 *big.Rat) *big.Rat {
	return new(bigRat).Mul(amount, amount2)
}

func bigRatDiv(amount *big.Rat, amount2 *big.Rat) *big.Rat {
	return new(bigRat).Quo(amount, amount2)
}
