package amountdecimal

import (
	"errors"
	"fmt"
)

// @title: initialization float32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewFloat32(amount float32) *AmountDecimal {
	amountStr := fmt.Sprintf("%f", amount)
	return NewString(amountStr)
}

// @title: addition float32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddFloat32(amounts ...float32) *AmountDecimal {
	return c.amountsFloat32(add, amounts...)
}

// @title: subtraction float32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubFloat32(amounts ...float32) *AmountDecimal {
	return c.amountsFloat32(sub, amounts...)
}

// @title: multiplication float32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulFloat32(amounts ...float32) *AmountDecimal {
	return c.amountsFloat32(mul, amounts...)
}

// @title: division float32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivFloat32(amounts ...float32) *AmountDecimal {
	return c.amountsFloat32(div, amounts...)
}

func (c *AmountDecimal) amountsFloat32(f uint8, amounts ...float32) *AmountDecimal {
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
