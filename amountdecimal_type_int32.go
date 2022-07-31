package amountdecimal

import "errors"

// @title: initialization int32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewInt32(amount int32) *AmountDecimal {
	return NewInt64(int64(amount))
}

// @title: addition int32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddInt32(amounts ...int32) *AmountDecimal {
	return c.amountsInt32(add, amounts...)
}

// @title: subtraction int32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubInt32(amounts ...int32) *AmountDecimal {
	return c.amountsInt32(sub, amounts...)
}

// @title: multiplication int32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulInt32(amounts ...int32) *AmountDecimal {
	return c.amountsInt32(mul, amounts...)
}

// @title: division int32 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivInt32(amounts ...int32) *AmountDecimal {
	return c.amountsInt32(div, amounts...)
}

func (c *AmountDecimal) amountsInt32(f uint8, amounts ...int32) *AmountDecimal {
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
