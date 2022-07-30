package amountdecimal

// @title: initialization uint8 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func NewUint8(amount uint8) *AmountDecimal {
	return NewUint64(uint64(amount))
}

// @title: addition uint8 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddUint8(amounts ...uint8) *AmountDecimal {
	return c.amountsUint8(add, amounts...)
}

// @title: subtraction uint8 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubUint8(amounts ...uint8) *AmountDecimal {
	return c.amountsUint8(sub, amounts...)
}

// @title: multiplication uint8 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulUint8(amounts ...uint8) *AmountDecimal {
	return c.amountsUint8(mul, amounts...)
}

// @title: division uint8 type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivUint8(amounts ...uint8) *AmountDecimal {
	return c.amountsUint8(div, amounts...)
}

func (c *AmountDecimal) amountsUint8(f uint8, amounts ...uint8) *AmountDecimal {
	var amounts2 []uint64
	for _, amount := range amounts {
		amounts2 = append(amounts2, uint64(amount))
	}

	return c.amountsUint64(f, amounts2...)
}