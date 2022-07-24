package amountdecimal

import (
	"math/big"
)

func NewBigRat(amount *big.Rat) *AmountDecimal {
	return NewString(amount.String())
}

func (c *AmountDecimal) AddBigRat(amount *big.Rat) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigRat(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubBigRat(amount *big.Rat) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigRat(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulBigRat(amount *big.Rat) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigRat(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivBigRat(amount *big.Rat) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewBigRat(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}