package amountdecimal

import (
	"errors"
	"math/big"
)

// @title: 金额校验
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func amountChk(amount string) (string, error) {
	if len(amount) == 0 {
		return amount, errors.New(errCodeMap[amount_not_numeric])
	}

	var isSpot bool
	l := len(amount) - 1
	anyReplaceList := []int32{'-', '.', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	for k, i := range amount {
		// 有空格
		if i == spaceInt32 {
			return amount, errors.New(errCodeMap[amount_not_numeric])
		}

		if i == spotInt32 {
			// 点号不能在末尾
			if k == l {
				return amount, errors.New(errCodeMap[amount_not_numeric])
			}

			// 至多有一个点号
			if isSpot {
				return amount, errors.New(errCodeMap[amount_not_numeric])
			}

			isSpot = true
		}

		// 减号要在第一位
		if i == negativeInt32 && k != 0 {
			return amount, errors.New(errCodeMap[amount_not_numeric])
		}

		if !int32InSlice(anyReplaceList, i) {
			return amount, errors.New(errCodeMap[amount_not_numeric])
		}
	}

	return amount, nil
}

// @title: int32数字是否在切片中
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func int32InSlice(list []int32, i int32) bool {
	for _, v := range list {
		if v == i {
			return true
		}
	}
	return false
}

// @title: 实例化big.Rat
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func newRat(amount string) (*big.Rat, error) {
	amountRat := new(big.Rat)
	_, ok := amountRat.SetString(amount)
	if !ok {
		return nil, errors.New(errCodeMap[bigrat_setstring_fail])
	}
	return amountRat, nil
}

// @title: 金额计算
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 22:43
func amountCalculation(f uint8, c *AmountDecimal, amount string) *AmountDecimal {
	var err error
	var data AmountDecimal

	if c.err != nil {
		return c
	}

	amount, err = amountChk(amount)
	if err != nil {
		c.err = err
		return c
	}

	amountRat, err := newRat(amount)
	if err != nil {
		c.err = err
		return c
	}

	switch f {
	case add:
		data.amount = new(big.Rat).Add(c.amount, amountRat)
	case sub:
		data.amount = new(big.Rat).Sub(c.amount, amountRat)
	case mul:
		data.amount = new(big.Rat).Mul(c.amount, amountRat)
	case div:
		data.amount = new(big.Rat).Quo(c.amount, amountRat)
	}

	data.decimal = c.decimal

	return &data
}
