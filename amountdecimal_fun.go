package amountdecimal

import (
	"encoding/json"
	"errors"
	"github.com/jishulangcom/go-fun"
	"math/big"
	"strconv"
)

func newAmountDecimal(amount interface{}) *AmountDecimal {
	var data AmountDecimal
	if amount == nil {
		data.err = errors.New(errCodeMap[amount_type_wrong])
		return &data
	}

	switch amount.(type) {
	case string:
		val, ok := amount.(string)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewString(val)
	case float64:
		val, ok := amount.(float64)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewFloat64(val)
	case float32:
		val, ok := amount.(float32)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewFloat32(val)
	case int:
		val, ok := amount.(int)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewInt(val)
	case int8:
		val, ok := amount.(int8)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewInt8(val)
	case int16:
		val, ok := amount.(int16)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewInt16(val)
	case int32:
		val, ok := amount.(int32)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewInt32(val)
	case int64:
		val, ok := amount.(int64)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewInt64(val)
	case uint:
		val, ok := amount.(uint)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewUint(val)
	case uint8:
		val, ok := amount.(uint8)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewUint8(val)
	case uint16:
		val, ok := amount.(uint16)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewUint16(val)
	case uint32:
		val, ok := amount.(uint32)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewUint32(val)
	case uint64:
		val, ok := amount.(uint64)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewUint64(val)
	case *big.Int:
		val, ok := amount.(*big.Int)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewBigInt(val)
	case *big.Rat:
		val, ok := amount.(*big.Rat)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewBigRat(val)

	default:
		data.err = errors.New(errCodeMap[amount_type_wrong])
		return &data
	}
}

// @title: 获取精度
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 23:15
func getDecimal(decimalOrCoin interface{}) (int, error) {
	decimalType := fun.GetType(decimalOrCoin)

	if decimalType == type_string {
		coin := InterfaceToStr(decimalOrCoin)
		coin = fun.Strtoupper(coin)
		if coinInfo, ok := CoinMap[coin]; ok {
			return coinInfo.Decimal, nil
		}

		return 0, errors.New("coin wrong")
	}

	if _, ok := numberTypeMap[decimalType]; ok {
		decimal, ok2 := decimalOrCoin.(int)
		if ok2{
			return decimal, nil
		}
		return 0, errors.New("decimal wrong")
	}



	return 0, errors.New("decimal type wrong")
}

func InterfaceToStr(val interface{}) string {
	var str string

	if val == nil {
		return str
	}

	switch val.(type) {
	case float64:
		ft := val.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := val.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := val.(int)
		str = strconv.Itoa(it)
	case uint:
		it := val.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := val.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := val.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := val.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := val.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := val.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := val.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := val.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := val.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = val.(string)
	case []byte:
		str = string(val.([]byte))
	default:
		newValue, _ := json.Marshal(val)
		str = string(newValue)
	}

	return str
}


/*
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
*/
