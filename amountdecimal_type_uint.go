package amountdecimal

//----------------------------------------------------------------------------------------------------------------------
func NewUint(amount uint) *AmountDecimal {
	return NewUint64(uint64(amount))
}

func NewUint64(amount uint64) *AmountDecimal {
	var data AmountDecimal

	data.amount = new(bigRat).SetUint64(amount)

	return &data
}

func NewUint32(amount uint32) *AmountDecimal {
	return NewUint64(uint64(amount))
}

func NewUint16(amount uint16) *AmountDecimal {
	return NewUint64(uint64(amount))
}

func NewUint8(amount uint8) *AmountDecimal {
	return NewUint64(uint64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) AddUint(amount uint) *AmountDecimal {
	return c.AddUint64(uint64(amount))
}

func (c *AmountDecimal) AddUint64(amount uint64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewUint64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) AddUint32(amount uint32) *AmountDecimal {
	return c.AddUint64(uint64(amount))
}

func (c *AmountDecimal) AddUint16(amount uint16) *AmountDecimal {
	return c.AddUint64(uint64(amount))
}

func (c *AmountDecimal) AddUint8(amount uint8) *AmountDecimal {
	return c.AddUint64(uint64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) SubUint(amount uint) *AmountDecimal {
	return c.SubInt64(int64(amount))
}

func (c *AmountDecimal) SubUint64(amount uint64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewUint64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubUint32(amount uint32) *AmountDecimal {
	return c.SubUint64(uint64(amount))
}

func (c *AmountDecimal) SubUint16(amount uint16) *AmountDecimal {
	return c.SubUint64(uint64(amount))
}

func (c *AmountDecimal) SubUint8(amount uint8) *AmountDecimal {
	return c.SubUint64(uint64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) MulUint(amount uint) *AmountDecimal {
	return c.MulUint64(uint64(amount))
}

func (c *AmountDecimal) MulUint64(amount uint64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewUint64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulUint32(amount uint32) *AmountDecimal {
	return c.MulUint64(uint64(amount))
}

func (c *AmountDecimal) MulUint16(amount uint16) *AmountDecimal {
	return c.MulUint64(uint64(amount))
}

func (c *AmountDecimal) MulUint8(amount uint8) *AmountDecimal {
	return c.MulUint64(uint64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) DivUint(amount uint) *AmountDecimal {
	return c.DivUint64(uint64(amount))
}

func (c *AmountDecimal) DivUint64(amount uint64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewUint64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivUint32(amount uint32) *AmountDecimal {
	return c.DivUint64(uint64(amount))
}

func (c *AmountDecimal) DivUint16(amount uint16) *AmountDecimal {
	return c.DivUint64(uint64(amount))
}

func (c *AmountDecimal) DivUint8(amount uint8) *AmountDecimal {
	return c.DivUint64(uint64(amount))
}
