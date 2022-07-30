package amountdecimal

// @title: initialization int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewInt(amount int) *AmountDecimal {
	return NewInt64(int64(amount))
}

// @title: addition int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddInt(amount int) *AmountDecimal {
	return c.AddInt64(int64(amount))
}

// @title: subtraction int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubInt(amount int) *AmountDecimal {
	return c.SubInt64(int64(amount))
}

// @title: multiplication int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulInt(amount int) *AmountDecimal {
	return c.MulInt64(int64(amount))
}

// @title: division int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivInt(amount int) *AmountDecimal {
	return c.DivInt64(int64(amount))
}