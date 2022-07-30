package amountdecimal

// @title: initialization int16 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewInt16(amount int16) *AmountDecimal {
	return NewInt64(int64(amount))
}

// @title: addition int16 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddInt16(amounts ...int16) *AmountDecimal {
	return c.amountsInt16(add, amounts...)
}

// @title: subtraction int16 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubInt16(amounts ...int16) *AmountDecimal {
	return c.amountsInt16(sub, amounts...)
}

// @title: multiplication int16 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulInt16(amounts ...int16) *AmountDecimal {
	return c.amountsInt16(mul, amounts...)
}

// @title: division int16 type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivInt16(amounts ...int16) *AmountDecimal {
	return c.amountsInt16(div, amounts...)
}

func (c *AmountDecimal) amountsInt16(f uint8, amounts ...int16) *AmountDecimal {
	var amounts2 []int64
	for _, amount := range amounts {
		amounts2 = append(amounts2, int64(amount))
	}

	return c.amountsInt64(f, amounts2...)
}
