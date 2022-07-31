package amountdecimal

import "errors"

// @title: initialization int8 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewInt8(amount int8) *AmountDecimal {
	return NewInt64(int64(amount))
}

// @title: addition int64 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddInt8(amounts ...int8) *AmountDecimal {
	return c.amountsInt8(add, amounts...)
}

// @title: subtraction int8 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubInt8(amounts ...int8) *AmountDecimal {
	return c.amountsInt8(sub, amounts...)
}

// @title: multiplication int8 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulInt8(amounts ...int8) *AmountDecimal {
	return c.amountsInt8(mul, amounts...)
}

// @title: division int8 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivInt8(amounts ...int8) *AmountDecimal {
	return c.amountsInt8(div, amounts...)
}

func (c *AmountDecimal) amountsInt8(f uint8, amounts ...int8) *AmountDecimal {
	if c.err != nil {
		return c
	}

	if len(amounts) == 0 {
		c.err = errors.New(errCodeMap[amounts_empty])
		return c
	}

	var amounts2 []int64
	for _, amount := range amounts {
		amounts2 = append(amounts2, int64(amount))
	}

	return c.amountsInt64(f, amounts2...)
}
