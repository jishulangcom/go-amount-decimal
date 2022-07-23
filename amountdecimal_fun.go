package amountdecimal

import (
	"errors"
	"fmt"
	"github.com/jishulangcom/go-fun"
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
	for k, i := range amount {
		// 有空格
		if i == int32_space {
			return amount, errors.New(errCodeMap[amount_not_numeric])
		}

		if i == int32_spot {
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
		if i == int32_negative && k != 0 {
			return amount, errors.New(errCodeMap[amount_not_numeric])
		}

		if !int32InSlice(amountElementList, i) {
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
	amountBigRat := new(big.Rat)
	_, ok := amountBigRat.SetString(amount)
	if !ok {
		return nil, errors.New(errCodeMap[bigrat_setstring_fail])
	}
	return amountBigRat, nil
}

// @title: 金额计算
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 22:43
func amountCalculation(f uint8, c *AmountDecimal, amount interface{}) *AmountDecimal {
	var err error
	var data AmountDecimal

	if c.err != nil {
		return c
	}

	amountBitRat, err := amountRat(amount)
	if err != nil {
		c.err = err
		return c
	}

	switch f {
	case add:
		data.amount = new(big.Rat).Add(c.amount, amountBitRat)
	case sub:
		data.amount = new(big.Rat).Sub(c.amount, amountBitRat)
	case mul:
		data.amount = new(big.Rat).Mul(c.amount, amountBitRat)
	case div:
		data.amount = new(big.Rat).Quo(c.amount, amountBitRat)
	}

	data.decimal = c.decimal

	return &data
}

func amountRat(amount interface{}) (*big.Rat, error) {
	if amount == nil {
		return nil, errors.New(errCodeMap[amount_type_wrong])
	}

	amountType := fun.GetType(amount)

	// 指针类型
	if amountType == type_ptr {
		// *big.Rat类型直接返回
		amountBitRat,ok := amount.(*big.Rat)
		if ok {
			return amountBitRat, nil
		}

		// *big.Int转string
		amountBitInt,ok := amount.(*big.Int)
		if ok {
			amountType = type_str
			amount = amountBitInt.String()
		} else {
			return nil, errors.New(errCodeMap[amount_type_wrong])
		}
	}

	amountStr, err := interfaceToString(amount, amountType)
	if err != nil {
		return nil, err
	}

	amountStr, err = amountChk(amountStr)
	if err != nil {
		return nil, err
	}

	amountBitRat, err := newRat(amountStr)
	if err != nil {
		return nil, err
	}
	return amountBitRat, nil
}

// @title: 金额计算
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 22:43
func interfaceToString (amount interface{}, amountType string) (string, error){
	amountStr := ""

	if amount == nil {
		return "", errors.New(errCodeMap[amount_type_wrong])
	}

	if amountType == type_str {
		amountStr = fmt.Sprintf("%s", amount)
		return amountStr, nil
	}

	if _, ok := numberTypeMap[amountType]; ok {
		amountStr = fmt.Sprintf("%v", amount)
		return amountStr, nil
	}

	fmt.Println("amountType:", amountType)
	return "", errors.New(errCodeMap[amount_type_wrong])
}