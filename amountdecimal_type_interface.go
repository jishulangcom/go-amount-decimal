/*
	amountdecimal
	【包名:】金额的高精度计算类库
	【作者:】jishulang.com
*/
package amountdecimal

import "errors"

// @title: initialization
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func New(amount interface{}) *AmountDecimal {
	return newAmountDecimal(amount)
}

// @title: addition
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Add(amounts ...interface{}) *AmountDecimal {
	return c.amounts(add, amounts...)
}

// @title: subtraction
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Sub(amounts ...interface{}) *AmountDecimal {
	return c.amounts(sub, amounts...)
}

// @title: multiplication
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Mul(amounts ...interface{}) *AmountDecimal {
	return c.amounts(mul, amounts...)
}

// @title: division
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Div(amounts ...interface{}) *AmountDecimal {
	return c.amounts(div, amounts...)
}

func (c *AmountDecimal) amounts(f uint8, amounts ...interface{}) *AmountDecimal {
	if c.err != nil {
		return c
	}

	if len(amounts) == 0 {
		c.err = errors.New(errCodeMap[amounts_empty])
		return c
	}

	var ad *AmountDecimal
	ad = c
	for _, amount := range amounts {
		ad = amountCalculation(f, ad, amount)
		if ad.err != nil {
			return ad
		}
	}
	return ad
}
