package amountdecimal

// @title: 实例化
// @param: amount string
// @param: decimal int
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func New(amount interface{}, decimal int) *AmountDecimal {
	var data AmountDecimal

	amountBitRat, err := amountRat(amount)
	if err != nil {
		data.err = err
		return &data
	}

	data.amount = amountBitRat
	data.decimal = decimal

	return &data
}

// @title: 加法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Add(amount interface{}) *AmountDecimal {
	return amountCalculation(add, c, amount)
}

// @title: 减法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Sub(amount interface{}) *AmountDecimal {
	return amountCalculation(sub, c, amount)
}

// @title: 乘法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Mul(amount interface{}) *AmountDecimal {
	return amountCalculation(mul, c, amount)
}

// @title: 除法
// @param: amount string
// @return: *AmountDecimal
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) Div(amount interface{}) *AmountDecimal {
	return amountCalculation(div, c, amount)
}
