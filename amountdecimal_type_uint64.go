package amountdecimal

import "errors"

// @title: initialization uint64 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func NewUint64(amount uint64) *AmountDecimal {
	var data AmountDecimal

	data.amount = new(bigRat).SetUint64(amount)

	return &data
}

// @title: addition uint64 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddUint64(amounts ...uint64) *AmountDecimal {
	return c.amountsUint64(add, amounts...)
}

// @title: subtraction uint64 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubUint64(amounts ...uint64) *AmountDecimal {
	return c.amountsUint64(sub, amounts...)
}

// @title: multiplication uint64 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulUint64(amounts ...uint64) *AmountDecimal {
	return c.amountsUint64(mul, amounts...)
}

// @title: division uint64 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivUint64(amounts ...uint64) *AmountDecimal {
	return c.amountsUint64(mul, amounts...)
}

func (c *AmountDecimal) amountsUint64(f uint8, amounts ...uint64) *AmountDecimal {
	if c.err != nil {
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

		amountDecimal := NewUint64(amount)
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
