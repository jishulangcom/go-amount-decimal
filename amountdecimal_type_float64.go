package amountdecimal

import (
	"errors"
	"fmt"
)

// @title: initialization float64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewFloat64(amount float64) *AmountDecimal {
	amountStr := fmt.Sprintf("%f", amount)
	return NewString(amountStr)
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

	var amounts2 []string
	amountStr := ""
	for _, amount := range amounts {
		if f == div && amount == 0 {
			c.err = errors.New(errCodeMap[amount_divisor_zero])
			return c
		}

		amountStr = fmt.Sprintf("%f", amount)
		amounts2 = append(amounts2, amountStr)
	}

	return c.amountsString(f, amounts2...)
}
