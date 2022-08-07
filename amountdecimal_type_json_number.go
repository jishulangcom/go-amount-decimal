package amountdecimal

import (
	"encoding/json"
	"errors"
	"fmt"
)

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

	if len(amounts) == 0 {
		c.err = errors.New(errCodeMap[amounts_empty])
		return c
	}

	var amounts2 []string
	amountStr := ""
	for _, amount := range amounts {
		if f == div && amount == json.Number(0) {
			c.err = errors.New(errCodeMap[amount_divisor_zero])
			return c
		}

		amountStr = fmt.Sprintf("%v", amount)
		amounts2 = append(amounts2, amountStr)
	}

	return c.amountsString(f, amounts2...)
}
