package amountdecimal

import "encoding/json"

func NewJsonNumber(amount json.Number) *AmountDecimal {
	return NewString(amount.String())
}

func (c *AmountDecimal) AddJsonNumber(amount json.Number) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount.String())
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubJsonNumber(amount json.Number) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount.String())
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulJsonNumber(amount json.Number) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount.String())
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivJsonNumber(amount json.Number) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount.String())
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}
