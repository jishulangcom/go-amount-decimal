package amountdecimal

import "math/big"

type AmountDecimal struct {
	amount  *big.Rat // 金额
	decimal int      // 精度
	err     error
}

const (
	add = iota // 加法
	sub        // 减法
	mul        // 乘法
	div        // 除法
)

const (
	negativeInt32 = '-' // 减号int32
	spotInt32     = '.' // 点号int32
	spaceInt32    = ' ' // 空格int32
)

const (
	bigrat_setstring_fail = iota
	bigint_setstring_fail
	string_fail
	amount_not_numeric
)

var errCodeMap map[uint16]string

func init() {
	errCodeMap = make(map[uint16]string)
	errCodeMap[bigrat_setstring_fail] = "big.Rat setString fail"
	errCodeMap[bigint_setstring_fail] = "big.Int setString fail"
	errCodeMap[string_fail] = "string fail"
	errCodeMap[amount_not_numeric] = "amount is not of numeric type"
}
