package amountdecimal

import (
	"errors"
	"math/big"
)

// @title: initialization string type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func NewString(amount string) *AmountDecimal {
	var data AmountDecimal

	amountBigRat, ok := new(big.Rat).SetString(amount)
	if !ok {
		data.err = errors.New(errCodeMap[string_setstring_fail])
	}

	data.amount = amountBigRat

	return &data
}

// @title: addition string type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddString(amounts ...string) *AmountDecimal {
	return c.amountsString(add, amounts...)
}

// @title: subtraction string type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubString(amounts ...string) *AmountDecimal {
	return c.amountsString(sub, amounts...)
}

// @title: multiplication string type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulString(amounts ...string) *AmountDecimal {
	return c.amountsString(mul, amounts...)
}

// @title: division string type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivString(amounts ...string) *AmountDecimal {
	return c.amountsString(div, amounts...)
}

func (c *AmountDecimal) amountsString(f uint8, amounts ...string) *AmountDecimal {
	if c.err != nil {
		return c
	}

	var ad *AmountDecimal
	ad = c
	for _, amount := range amounts {
		amountDecimal := NewString(amount)
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
