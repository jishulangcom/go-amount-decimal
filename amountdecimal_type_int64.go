package amountdecimal

import "errors"

// @title: initialization int64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewInt64(amount int64) *AmountDecimal {
	var data AmountDecimal

	data.amount = new(bigRat).SetInt64((amount))

	return &data
}

// @title: addition int64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddInt64(amounts ...int64) *AmountDecimal {
	return c.amountsInt64(add, amounts...)
}

// @title: subtraction int64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubInt64(amounts ...int64) *AmountDecimal {
	return c.amountsInt64(sub, amounts...)
}

// @title: multiplication int64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulInt64(amounts ...int64) *AmountDecimal {
	return c.amountsInt64(mul, amounts...)
}

// @title: division int64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivInt64(amounts ...int64) *AmountDecimal {
	return c.amountsInt64(div, amounts...)
}

func (c *AmountDecimal) amountsInt64(f uint8, amounts ...int64) *AmountDecimal {
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
		if f == div && ad.amount.String() == bigrat_zero_string {
			return ad
		}

		if f == div && amount == 0 {
			ad.err = errors.New(errCodeMap[amount_divisor_zero])
			return ad
		}

		amountDecimal := NewInt64(amount)
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
