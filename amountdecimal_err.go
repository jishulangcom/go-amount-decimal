package amountdecimal


// errorcode
const (
	bigrat_setstring_fail = iota
	bigint_setstring_fail
	string_setstring_fail
	to_string_panic
	amount_not_numeric
	amount_empty
	amounts_empty
	amount_type_wrong
	amount_type_conversion
	amount_divisor_cannot
	coin_wrong
	coin_type_wrong
)

var errCodeMap map[uint16]string

func init() {
	errCodeMap = make(map[uint16]string)
	errCodeMap[bigrat_setstring_fail] = "big.Rat setString fail"
	errCodeMap[bigint_setstring_fail] = "big.Int setString fail"
	errCodeMap[string_setstring_fail] = "string setString fail"

	errCodeMap[to_string_panic] = "ToString() panic"
	errCodeMap[amount_empty] = "amount empty"
	errCodeMap[amounts_empty] = "amounts empty"
	errCodeMap[amount_not_numeric] = "amount is not of numeric type"
	errCodeMap[amount_type_wrong] = "amount type is not supported"
	errCodeMap[amount_type_conversion] = "amount type conversion failed"
	errCodeMap[amount_divisor_cannot] = "divisor cannot be 0"
	errCodeMap[coin_wrong] = "coin wrong"
	errCodeMap[coin_type_wrong] = "decimal type wrong"
}
