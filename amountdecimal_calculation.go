package amountdecimal

import (
	"errors"
	"math/big"
)

// @title: 金额计算
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 22:43
func amountCalculation(f uint8, c *AmountDecimal, amount interface{}) *AmountDecimal {
	var err error

	if c.err != nil {
		return c
	}

	data := newAmountDecimal(amount)
	if err != nil {
		return data
	}

	return calculationBigRat(f, c.amount, data.amount)
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func calculationBigRat(f uint8, amount *big.Rat, amount2 *big.Rat) *AmountDecimal {
	var data AmountDecimal

	// 计算方法所对应的操作
	switch f {
	case add:
		data.amount = new(bigRat).Add(amount, amount2)
	case sub:
		data.amount = new(bigRat).Sub(amount, amount2)
	case mul:
		data.amount = new(bigRat).Mul(amount, amount2)
	case div:
		if amount.String() == "0/1" {
			data.err = errors.New(errCodeMap[amount_divisor_cannot])
			return &data
		}
		data.amount = new(bigRat).Quo(amount, amount2)
	}

	return &data
}
