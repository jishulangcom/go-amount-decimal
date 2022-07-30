/*
	amountdecimal
	【包名:】金额的高精度计算类库
	【作者:】技术狼(jishulang.com)
*/
package amountdecimal

// @title: 实例化
// @param: amount string
// @param: decimal int
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func New(amount interface{}) *AmountDecimal {
	return newAmountDecimal(amount)
}

// @title: 加法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Add(amounts ...interface{}) *AmountDecimal {
	for _, amount := range amounts {
		c = amountCalculation(add, c, amount)
		if c.err != nil {
			return c
		}
	}
	return c
}

// @title: 减法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Sub(amounts ...interface{}) *AmountDecimal {
	for _, amount := range amounts {
		c = amountCalculation(sub, c, amount)
		if c.err != nil {
			return c
		}
	}
	return c
}

// @title: 乘法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Mul(amounts ...interface{}) *AmountDecimal {
	for _, amount := range amounts {
		c = amountCalculation(mul, c, amount)
		if c.err != nil {
			return c
		}
	}
	return c
}

// @title: 除法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Div(amounts ...interface{}) *AmountDecimal {
	for _, amount := range amounts {
		c = amountCalculation(div, c, amount)
		if c.err != nil {
			return c
		}
	}
	return c
}
