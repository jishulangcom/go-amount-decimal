package amountdecimal

import (
	"math/big"
)

type bigRat struct {
	big.Rat
}

const bigrat_zero_string = "0/1"

var bigrat_zero *big.Rat
var bigint_zero *big.Int

const (
	int32_negative = '-' // 减号int32
	int32_spot     = '.' // 点号int32
	int32_space    = ' ' // 空格int32
)

// 数据类型
const (
	type_ptr     = "ptr"
	type_string  = "string"
	type_float32 = "float32"
	type_float64 = "float64"
)

// errorcode
const (
	bigrat_setstring_fail = iota
	bigint_setstring_fail
	string_setstring_fail
	to_string_panic
	amount_not_numeric
	amount_empty
	amount_type_wrong
	amount_type_conversion
	amount_divisor_cannot
	coin_wrong
	coin_type_wrong
)

var DefaultDecimal = 13 // 默认输出精度

var errCodeMap map[uint16]string
var amountElementList []int32         // 金额组成元素
var numberTypeMap map[string]struct{} // 数据类型集合

func init() {
	errCodeMap = make(map[uint16]string)
	errCodeMap[bigrat_setstring_fail] = "big.Rat setString fail"
	errCodeMap[bigint_setstring_fail] = "big.Int setString fail"
	errCodeMap[string_setstring_fail] = "string setString fail"

	errCodeMap[to_string_panic] = "ToString() panic"
	errCodeMap[amount_empty] = "amount not empty"
	errCodeMap[amount_not_numeric] = "amount is not of numeric type"
	errCodeMap[amount_type_wrong] = "amount type is not supported"
	errCodeMap[amount_type_conversion] = "amount type conversion failed"
	errCodeMap[amount_divisor_cannot] = "divisor cannot be 0"
	errCodeMap[coin_wrong] = "coin wrong"
	errCodeMap[coin_type_wrong] = "decimal type wrong"

	amountElementList = []int32{'e', '+', '-', '.', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	list := []string{
		//"float32", "float64",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
	}
	numberTypeMap = make(map[string]struct{})
	for _, v := range list {
		numberTypeMap[v] = struct{}{}
	}

	bigrat_zero = new(bigRat).SetInt64(0)
	bigint_zero = new(big.Int).SetInt64(0)
}
