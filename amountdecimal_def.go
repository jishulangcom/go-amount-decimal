package amountdecimal

import (
	"math/big"
)

type bigRat struct {
	big.Rat
}

// const bigrat_zero_string = "0/1"

var bigrat_zero *big.Rat
var bigint_zero *big.Int

// 数据类型
const (
	// type_ptr     = "ptr"
	type_string = "string"
	// type_float32 = "float32"
	// type_float64 = "float64"
)

var DefaultDecimal = 13               // 默认输出精度
var numberTypeMap map[string]struct{} // 数据类型集合

func init() {
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
