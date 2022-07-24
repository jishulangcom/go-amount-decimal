package amountdecimal

import "errors"

//----------------------------------------------------------------------------------------------------------------------
func NewInt(amount int) *AmountDecimal {
	return NewInt64(int64(amount))
}

func NewInt64(amount int64) *AmountDecimal {
	var data AmountDecimal

	data.amount = new(bigRat).SetInt64((amount))

	return &data
}

func NewInt32(amount int32) *AmountDecimal {
	return NewInt64(int64(amount))
}

func NewInt16(amount int16) *AmountDecimal {
	return NewInt64(int64(amount))
}

func NewInt8(amount int8) *AmountDecimal {
	return NewInt64(int64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) AddInt(amount int) *AmountDecimal {
	return c.AddInt64(int64(amount))
}

func (c *AmountDecimal) AddInt64(amount int64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewInt64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) AddInt32(amount int32) *AmountDecimal {
	return c.AddInt64(int64(amount))
}

func (c *AmountDecimal) AddInt16(amount int16) *AmountDecimal {
	return c.AddInt64(int64(amount))
}

func (c *AmountDecimal) AddInt8(amount int8) *AmountDecimal {
	return c.AddInt64(int64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) SubInt(amount int) *AmountDecimal {
	return c.SubInt64(int64(amount))
}

func (c *AmountDecimal) SubInt64(amount int64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewInt64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubInt32(amount int32) *AmountDecimal {
	return c.SubInt64(int64(amount))
}

func (c *AmountDecimal) SubInt16(amount int16) *AmountDecimal {
	return c.SubInt64(int64(amount))
}

func (c *AmountDecimal) SubInt8(amount int8) *AmountDecimal {
	return c.SubInt64(int64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) MulInt(amount int) *AmountDecimal {
	return c.MulInt64(int64(amount))
}

func (c *AmountDecimal) MulInt64(amount int64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewInt64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulInt32(amount int32) *AmountDecimal {
	return c.MulInt64(int64(amount))
}

func (c *AmountDecimal) MulInt16(amount int16) *AmountDecimal {
	return c.MulInt64(int64(amount))
}

func (c *AmountDecimal) MulInt8(amount int8) *AmountDecimal {
	return c.MulInt64(int64(amount))
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) DivInt(amount int) *AmountDecimal {
	return c.DivInt64(int64(amount))
}

func (c *AmountDecimal) DivInt64(amount int64) *AmountDecimal {
	if c.err != nil {
		return c
	}

	if amount == 0 {
		c.err = errors.New(errCodeMap[amount_divisor_cannot])
		return c
	}

	amountDecimal := NewInt64(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivInt32(amount int32) *AmountDecimal {
	return c.DivInt64(int64(amount))
}

func (c *AmountDecimal) DivInt16(amount int16) *AmountDecimal {
	return c.DivInt64(int64(amount))
}

func (c *AmountDecimal) DivInt8(amount int8) *AmountDecimal {
	return c.DivInt64(int64(amount))
}
