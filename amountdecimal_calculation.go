package amountdecimal

import "math/big"

// @title: 金额计算
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 22:43
func amountCalculation(f uint8, c *AmountDecimal, amount interface{}) *AmountDecimal {
	var err error

	if c.err != nil {
		return c
	}

	amountBigRat, err := amountRat(amount)
	if err != nil {
		c.err = err
		return c
	}

	return calculationBigRat(f, c.amount, amountBigRat)
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
		data.amount = new(bigRat).Quo(amount, amount2)
	}

	return &data
}
