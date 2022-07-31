package amountdecimal

import "errors"

// @title: initialization float64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewFloat64(amount float64) *AmountDecimal {
	var data AmountDecimal

	data.decimal = float64Decimal(amount)
	data.amount = new(bigRat).SetFloat64(amount)

	return &data
}

// @title: addition float64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddFloat64(amounts ...float64) *AmountDecimal {
	return c.amountsFloat64(add, amounts...)
}

// @title: subtraction float64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubFloat64(amounts ...float64) *AmountDecimal {
	return c.amountsFloat64(sub, amounts...)
}

// @title: multiplication float64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulFloat64(amounts ...float64) *AmountDecimal {
	return c.amountsFloat64(mul, amounts...)
}

// @title: division float64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivFloat64(amounts ...float64) *AmountDecimal {
	return c.amountsFloat64(div, amounts...)
}

func (c *AmountDecimal) amountsFloat64(f uint8, amounts ...float64) *AmountDecimal {
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

		if f == div && amount == 0 {
			ad.err = errors.New(errCodeMap[amount_divisor_cannot])
			return ad
		}

		amountDecimal := NewFloat64(amount)
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
