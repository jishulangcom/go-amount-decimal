package amountdecimal

// @title: initialization uint32 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func NewUint32(amount uint32) *AmountDecimal {
	return NewUint64(uint64(amount))
}

// @title: addition uint32 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddUint32(amounts ...uint32) *AmountDecimal {
	return c.amountsUint32(add, amounts...)
}

// @title: subtraction uint32 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubUint32(amounts ...uint32) *AmountDecimal {
	return c.amountsUint32(sub, amounts...)
}

// @title: multiplication uint8 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulUint32(amounts ...uint32) *AmountDecimal {
	return c.amountsUint32(mul, amounts...)
}

// @title: division uint32 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivUint32(amounts ...uint32) *AmountDecimal {
	return c.amountsUint32(div, amounts...)
}

func (c *AmountDecimal) amountsUint32(f uint8, amounts ...uint32) *AmountDecimal {
	var amounts2 []uint64
	for _, amount := range amounts {
		amounts2 = append(amounts2, uint64(amount))
	}

	return c.amountsUint64(f, amounts2...)
}