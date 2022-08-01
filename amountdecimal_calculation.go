package amountdecimal

import (
	"encoding/json"
	"errors"
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

	switch f {
	case add:
		if c.decimal > ad.decimal {
			data.decimal = c.decimal
		} else {
			data.decimal = ad.decimal
		}
	case sub:
		if c.decimal > ad.decimal {
			data.decimal = c.decimal
		} else {
			data.decimal = ad.decimal
		}
	case mul:
		data.decimal = c.decimal + ad.decimal
	case div:
		data.decimal = DefaultDecimal

	}

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

	switch f {
	case add:
		data.amount = new(bigRat).Add(amount, amount2)
	case sub:
		data.amount = new(bigRat).Sub(amount, amount2)
	case mul:
		data.amount = new(bigRat).Mul(amount, amount2)
	case div:
		// dividend is 0
		if amount == bigrat_zero {
			data.amount = amount
			return &data
		}

		// divide is 0
		if amount2 == bigrat_zero {
			data.err = errors.New(errCodeMap[amount_divisor_cannot])
			return &data
		}

		data.amount = new(bigRat).Quo(amount, amount2)
	}

	return &data
}
