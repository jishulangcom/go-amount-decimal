package amountdecimal

import (
	"errors"
	"math/big"
)

// @title: initialization *big.Rat type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewBigRat(amount *big.Rat) *AmountDecimal {
	return NewString(amount.String())
}

// @title: addition *big.Rat type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddBigRat(amounts ...*big.Rat) *AmountDecimal {
	return c.amountsBigRat(add, amounts...)
}

// @title: subtraction *big.Rat type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubBigRat(amounts ...*big.Rat) *AmountDecimal {
	return c.amountsBigRat(sub, amounts...)
}

// @title: multiplication *big.Rat type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulBigRat(amounts ...*big.Rat) *AmountDecimal {
	return c.amountsBigRat(mul, amounts...)
}

// @title: division *big.Rat type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivBigRat(amounts ...*big.Rat) *AmountDecimal {
	return c.amountsBigRat(div, amounts...)
}

func (c *AmountDecimal) amountsBigRat(f uint8, amounts ...*big.Rat) *AmountDecimal {
	if c.err != nil {
		return c
	}

	if len(amounts) == 0 {
		c.err = errors.New(errCodeMap[amounts_empty])
		return c
	}

	var ad *AmountDecimal
	ad = c
	for _, amount := range amounts {
		if f == div && ad.amount == bigrat_zero {
			return ad
		}

		if f == div && amount == bigrat_zero {
			ad.err = errors.New(errCodeMap[amount_divisor_cannot])
			return ad
		}

		amountDecimal := NewBigRat(amount)
		if amountDecimal.err != nil {
			return amountDecimal
		}

		ad = bigRatCalculation(f, ad.amount, amountDecimal.amount)
		if ad.err != nil {
			return ad
		}
	}

	return ad
}
