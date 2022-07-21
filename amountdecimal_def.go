package amountdecimal

import "math/big"

type AmountDecimal struct {
	amount  *big.Rat // 金额
	decimal int      // 精度
	err     error
}

const (
	add = iota
	sub
	mul
	div
)

const (
	bigrat_setstring_fail = iota
	bigint_setstring_fail
)

var errCodeMap map[uint16]string

func init()  {
	errCodeMap = make(map[uint16]string)
	errCodeMap[bigrat_setstring_fail] = "big.Rat setString fail"
	errCodeMap[bigint_setstring_fail] = "big.Int setString fail"
}
