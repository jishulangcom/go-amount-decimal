package amountdecimal

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

// @title: Amount Calculation
// @auth: jishulang.com
// @date: 2022/7/21 22:43
func amountCalculation(f uint8, c *AmountDecimal, amount interface{}) *AmountDecimal {
	if c.err != nil {
		return c
	}

	ad := newAmountDecimal(amount)
	if ad.err != nil {
		return ad
	}

	data := bigRatCalculation(f, c.amount, ad.amount)

	return data
}

// @title: new AmountDecimal
// @auth: jishulang.com
// @date: 2022/7/30 19:49
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
		val2 := fmt.Sprintf("%f", amount)
		return NewString(val2)
	case float32:
		val := fmt.Sprintf("%f", amount)
		return NewString(val)
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
	case json.Number:
		val, ok := amount.(json.Number)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewJsonNumber(val)
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
	case *AmountDecimal:
		val, ok := amount.(*AmountDecimal)
		if !ok {
			data.err = errors.New(errCodeMap[amount_type_conversion])
			return &data
		}
		return NewBigRat(val.amount)

	default:

	}

	data.err = errors.New(errCodeMap[amount_type_wrong])
	return &data
}

// @title: *big.Rat Calculation
// @auth: jishulang.com
// @date: 2022/7/30 19:49
func bigRatCalculation(f uint8, amount *big.Rat, amount2 *big.Rat) *AmountDecimal {
	var data AmountDecimal

	// dividend is 0
	//if amount == bigrat_zero {
	if f == div && amount.String() == bigrat_zero_string {
		data.amount = amount
		return &data
	}

	// divide is 0
	//if amount2 == bigrat_zero {
	if f == div && amount2.String() == bigrat_zero_string {
		data.err = errors.New(errCodeMap[amount_divisor_cannot])
		return &data
	}

	data.amount = bigRatFunMap[f](amount, amount2)
	return &data
}
