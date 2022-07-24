package amountdecimal

import (
	"errors"
	"fmt"
	"github.com/jishulangcom/go-fun"
	"math/big"
	"strconv"
)

// @title: 获取amount的*big.Rat
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func amountRat(amount interface{}) (*big.Rat, error) {
	if amount == nil {
		return nil, errors.New(errCodeMap[amount_type_wrong])
	}

	amountType := fun.GetType(amount)

	// 指针类型
	if amountType == type_ptr {
		// *big.Rat类型直接返回
		amountBigRat,ok := amount.(*big.Rat)
		if ok {
			return amountBigRat, nil
		}

		// *big.Int转string
		amountBitInt,ok := amount.(*big.Int)
		if ok {
			amountType = type_string
			amount = amountBitInt.String()
		} else {
			return nil, errors.New(errCodeMap[amount_type_wrong])
		}
	}

	amountStr, err := interfaceToString(amount, amountType)
	if err != nil {
		return nil, err
	}

	if amountType == type_string {
		amountStr, err = amountStrChk(amountStr)
		if err != nil {
			return nil, err
		}
	}

	amountBigRat, err := newBigRat(amountStr)
	if err != nil {
		return nil, err
	}
	return amountBigRat, nil
}

// @title: 金额校验
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func amountStrChk(amount string) (string, error) {
	if len(amount) == 0 {
		return amount, errors.New(errCodeMap[amount_empty])
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
func newBigRat(amount string) (*big.Rat, error) {
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
func interfaceToString (amount interface{}, amountType string) (string, error){
	amountStr := ""

	if amount == nil {
		return "", errors.New(errCodeMap[amount_type_wrong])
	}

	if amountType == type_string {
		amountStr = fmt.Sprintf("%s", amount)
		return amountStr, nil
	}

	if amountType == type_float32 || amountType == type_float64 {
		//amountStr = fmt.Sprintf("%v", amount) // 3499504500  3469279200
		ft := amount.(float64)
		amountStr = strconv.FormatFloat(ft, 'f', -1, 64)
		//fmt.Println("amountStr ff:", amountStr)
		return amountStr, nil
	}

	if _, ok := numberTypeMap[amountType]; ok {
		amountStr = fmt.Sprintf("%d", amount)
		return amountStr, nil
	}

	fmt.Println("amountType:", amountType)
	return "", errors.New(errCodeMap[amount_type_wrong])
}

