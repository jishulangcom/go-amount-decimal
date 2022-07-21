package amountdecimal

import (
	"errors"
	"math/big"
)

// @title: 实例化检测
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func newChk(amount string, decimal int) error {
	return nil
}

// @title: 实例化检测
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func amountChk(amount string) error {
	return nil
}

func newRat(amount string) (*big.Rat, error) {
	amountRat := new(big.Rat)
	_, ok := amountRat.SetString(amount)
	if !ok {
		return nil, errors.New(errCodeMap[bigrat_setstring_fail])
	}
	return amountRat, nil
}

// @title: 金额计算
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 22:43
func amountCalculation(f uint8, c *AmountDecimal, amount string) *AmountDecimal{
	var data AmountDecimal

	if c.err != nil {
		return c
	}

	amountRat, err := newRat(amount)
	if err != nil {
		c.err = err
		return c
	}

	switch f {
		case add:
			data.amount = new(big.Rat).Add(c.amount, amountRat)
		case sub:
			data.amount = new(big.Rat).Sub(c.amount, amountRat)
		case mul:
			data.amount = new(big.Rat).Sub(c.amount, amountRat)
	}

	data.decimal = c.decimal

	return &data
}