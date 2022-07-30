package amountdecimal

// @title: initialization uint16 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func NewUint16(amount uint16) *AmountDecimal {
	return NewUint64(uint64(amount))
}

// @title: addition uint16 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddUint16(amounts ...uint16) *AmountDecimal {
	return c.amountsUint16(mul, amounts...)
}

// @title: subtraction uint16 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubUint16(amounts ...uint16) *AmountDecimal {
	return c.amountsUint16(mul, amounts...)
}

// @title: multiplication uint16 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulUint16(amounts ...uint16) *AmountDecimal {
	return c.amountsUint16(mul, amounts...)
}

// @title: division uint16 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivUint16(amounts ...uint16) *AmountDecimal {
	return c.amountsUint16(div, amounts...)
}

func (c *AmountDecimal) amountsUint16(f uint8, amounts ...uint16) *AmountDecimal {
	var amounts2 []uint64
	for _, amount := range amounts {
		amounts2 = append(amounts2, uint64(amount))
	}

	return c.amountsUint64(f, amounts2...)
}
