package amountdecimal

import "errors"

// @title: initialization int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewInt(amount int) *AmountDecimal {
	return NewInt64(int64(amount))
}

// @title: addition int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddInt(amounts ...int) *AmountDecimal {
	return c.amountsInt(add, amounts...)
}

// @title: subtraction int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubInt(amounts ...int) *AmountDecimal {
	return c.amountsInt(sub, amounts...)
}

// @title: multiplication int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulInt(amounts ...int) *AmountDecimal {
	return c.amountsInt(mul, amounts...)
}

// @title: division int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivInt(amounts ...int) *AmountDecimal {
	return c.amountsInt(div, amounts...)
}

func (c *AmountDecimal) amountsInt(f uint8, amounts ...int) *AmountDecimal {
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
