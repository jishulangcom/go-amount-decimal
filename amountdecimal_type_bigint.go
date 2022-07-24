package amountdecimal

import (
	"fmt"
	"math/big"
)

func NewBigInt(amount *big.Int) *AmountDecimal {
	return NewString(amount.String())
}

func (c *AmountDecimal) AddBigInt(amount *big.Int) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigInt(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubBigInt(amount *big.Int) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigInt(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulBigInt(amount *big.Int) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigInt(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivBigInt(amount *big.Int) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigInt(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}


// @title: 获取*big.Int的实际金额
// @param: amount *big.Int
// @param: decimal int
// @return: string, error
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func BigIntActualAmount(amount *big.Int, decimal int) (string, error) {
	ad := NewBigInt(amount)
	if ad.err != nil {
		return "", ad.err
	}

	expStr := fmt.Sprintf("1e-%d", decimal)
	expAd := NewString(expStr)
	if expAd.err != nil {
		return "", expAd.err
	}

	result := new(bigRat).Mul(ad.amount, expAd.amount)
	return result.FloatString(decimal), nil
}



