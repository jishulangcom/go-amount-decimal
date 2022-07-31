package amountdecimal

import "errors"

// @title: initialization uint type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func NewUint(amount uint) *AmountDecimal {
	return NewUint64(uint64(amount))
}

// @title: addition uint type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) AddUint(amounts ...uint) *AmountDecimal {
	return c.amountsUint(add, amounts...)
}

// @title: subtraction uint type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) SubUint(amounts ...uint) *AmountDecimal {
	return c.amountsUint(sub, amounts...)
}

// @title: multiplication uint type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) MulUint(amounts ...uint) *AmountDecimal {
	return c.amountsUint(mul, amounts...)
}

// @title: division uint type
// @auth: jishulang.com
// @date: 2022/7/30 20:59
func (c *AmountDecimal) DivUint(amounts ...uint) *AmountDecimal {
	return c.amountsUint(div, amounts...)
}

func (c *AmountDecimal) amountsUint(f uint8, amounts ...uint) *AmountDecimal {
	if c.err != nil {
		return c
	}

	if len(amounts) == 0 {
		c.err = errors.New(errCodeMap[amounts_empty])
		return c
	}

	var amounts2 []uint64
	for _, amount := range amounts {
		amounts2 = append(amounts2, uint64(amount))
	}

	return c.amountsUint64(f, amounts2...)
}
