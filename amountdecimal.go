package amountdecimal

import (
	"math/big"
)

// @title: 实例化
// @param: amount string
// @param: decimal int
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func New(amount string, decimal int) *AmountDecimal {
	var data AmountDecimal

	err := newChk(amount, decimal)
	if err != nil {
		data.err = err
		return &data
	}

	amountRat, err := newRat(amount)
	if err != nil {
		data.err = err
		return &data
	}

	data.amount = amountRat
	data.decimal = decimal

	return &data
}

// @title: 加法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Add(amount string) *AmountDecimal {
	return amountCalculation(add, c, amount)
}

// @title: 减法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Sub(amount string) *AmountDecimal {
	return amountCalculation(sub, c, amount)
}

// @title: 乘法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Mul(amount string) *AmountDecimal {
	return amountCalculation(mul, c, amount)
}

// @title: 除法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Div(amount string) *AmountDecimal {
	var data AmountDecimal

	if c.err != nil {
		return c
	}

	amountRat, err := newRat(amount)
	if err != nil {
		c.err = err
		return c
	}

	data.amount = new(big.Rat).Mul(c.amount, amountRat)
	data.decimal = c.decimal

	return &data
}
