package amountdecimal

// errorcode
const (
	bigrat_setstring_fail = iota
	bigint_setstring_fail
	string_setstring_fail
	to_string_panic
	amount_not_numeric
	amounts_empty
	amount_type_wrong
	amount_type_conversion
	amount_divisor_zero
	coin_wrong
	coin_type_wrong
	decimalorcoin_required
)

var errCodeMap map[uint16]string

func init() {
	errCodeMap = map[uint16]string{
		bigrat_setstring_fail: "big.Rat setString fail",
		bigint_setstring_fail: "big.Int setString fail",
		string_setstring_fail: "string setString fail",

		to_string_panic: "ToString() panic",

		amounts_empty:          "amounts empty",
		amount_not_numeric:     "amount not numeric type",
		amount_type_wrong:      "amount type not supported",
		amount_type_conversion: "amount type conversion failed",
		amount_divisor_zero:    "divisor cannot be 0",

		coin_wrong:      "coin wrong",
		coin_type_wrong: "decimal type wrong",

		decimalorcoin_required: "decimalOrCoin required",
	}
}
