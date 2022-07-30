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
	var data *AmountDecimal

	if c.err != nil {
		return c
	}

	amountD := newAmountDecimal(amount)
	if err != nil {
		return amountD
	}

	data = calculationBigRat(f, c.amount, amountD.amount)

	// 小数点精度取最多的那个
	if f == mul {
		data.decimal = c.decimal + amountD.decimal
	} else if f == div {
		data.decimal = DefaultDecimal
	} else {
		if c.decimal > amountD.decimal {
			data.decimal = c.decimal
		} else {
			data.decimal = amountD.decimal
		}
	}

	return data
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
		// 除数为0
		if amount.String() == "0/1" {
			data.amount = amount
			return &data
		}

		// 数除为0
		if amount2.String() == "0/1" {
			data.err = errors.New(errCodeMap[amount_divisor_cannot])
			return &data
		}
		data.amount = new(bigRat).Quo(amount, amount2)
	}

	return &data
}
