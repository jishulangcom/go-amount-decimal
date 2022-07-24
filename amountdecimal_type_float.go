package amountdecimal

//----------------------------------------------------------------------------------------------------------------------
func NewFloat64(amount float64) *AmountDecimal {
	var data AmountDecimal

	data.amount = new(bigRat).SetFloat64(amount)

	return &data
}

func NewFloat32(amount float32) *AmountDecimal {
	return NewFloat64(float64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) AddFloat64(amount float64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewFloat64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) AddFloat32(amount float32) *AmountDecimal {
	return c.AddFloat64(float64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) SubFloat64(amount float64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewFloat64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubFloat32(amount float32) *AmountDecimal {
	return c.SubFloat64(float64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) MulFloat64(amount float64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewFloat64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulFloat32(amount float32) *AmountDecimal {
	return c.MulFloat64(float64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) DivFloat64(amount float64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewFloat64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivFloat32(amount float32) *AmountDecimal {
	return c.DivFloat64(float64(amount))
}