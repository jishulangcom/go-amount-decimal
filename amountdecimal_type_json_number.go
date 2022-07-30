package amountdecimal

import "encoding/json"

func NewJsonNumber(amount json.Number) *AmountDecimal {
	return NewString(amount.String())
}

// @title: addition json.Number type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddJsonNumber(amounts ...json.Number) *AmountDecimal {
	return c.amountsJsonNumber(add, amounts...)
}

// @title: subtraction json.Number type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubJsonNumber(amounts ...json.Number) *AmountDecimal {
	return c.amountsJsonNumber(sub, amounts...)
}

// @title: multiplication json.Number type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulJsonNumber(amounts ...json.Number) *AmountDecimal {
	return c.amountsJsonNumber(mul, amounts...)
}

// @title: division json.Number type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivJsonNumber(amounts ...json.Number) *AmountDecimal {
	return c.amountsJsonNumber(div, amounts...)
}

func (c *AmountDecimal) amountsJsonNumber(f uint8, amounts ...json.Number) *AmountDecimal {
	if c.err != nil {
		return c
	}

	var ad *AmountDecimal
	ad = c
	for _, amount := range amounts {
		amountDecimal := NewString(amount.String())
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
