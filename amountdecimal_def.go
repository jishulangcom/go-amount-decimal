package amountdecimal

import (
	"math/big"
)

type AmountDecimal struct {
	amount  *big.Rat // 金额
	err     error
}

type bigRat struct{
	big.Rat
}

// AmountDecimal 计算方法
const (
	add = iota // 加法
	sub        // 减法
	mul        // 乘法
	div        // 除法
)

const (
	int32_negative = '-' // 减号int32
	int32_spot     = '.' // 点号int32
	int32_space    = ' ' // 空格int32
)

// 数据类型
const (
	type_ptr = "ptr"
	type_string = "string"
	type_float32 = "float32"
	type_float64 = "float64"
)

// errorcode
const (
	bigrat_setstring_fail = iota
	bigint_setstring_fail
	string_fail
	amount_not_numeric
	amount_empty
	amount_type_wrong
)

var errCodeMap map[uint16]string
var amountElementList []int32         // 金额组成元素
var numberTypeMap map[string]struct{} // 数据类型集合

func init() {
	errCodeMap = make(map[uint16]string)
	errCodeMap[bigrat_setstring_fail] = "big.Rat setString fail"
	errCodeMap[bigint_setstring_fail] = "big.Int setString fail"
	errCodeMap[string_fail] = "string fail"
	errCodeMap[amount_empty] = "amount not empty"
	errCodeMap[amount_not_numeric] = "amount is not of numeric type"
	errCodeMap[amount_type_wrong] = "amount type is not supported"

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
}
