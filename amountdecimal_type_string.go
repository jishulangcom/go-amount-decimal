package amountdecimal

import (
	"errors"
	"math/big"
)

//----------------------------------------------------------------------------------------------------------------------
func NewString(amount string) *AmountDecimal {
	var data AmountDecimal

	amountBigRat, ok := new(big.Rat).SetString(amount)
	if !ok {
		data.err = errors.New("xxxxxxxxxxxxxxxxxxx")
	}

	data.amount = amountBigRat

	return &data
}

//----------------------------------------------------------------------------------------------------------------------
func (c *AmountDecimal) AddString(amount string) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(add, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) SubString(amount string) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(sub, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) MulString(amount string) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(mul, c.amount, amountDecimal.amount)
}

func (c *AmountDecimal) DivString(amount string) *AmountDecimal {
	if c.err != nil {
		return c
	}

	amountDecimal := NewString(amount)
	if amountDecimal.err != nil {
		return amountDecimal
	}

	return calculationBigRat(div, c.amount, amountDecimal.amount)
}