package amountdecimal

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

// @title: Output string
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) ToString(decimalOrCoin interface{}) (amountStr string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(errCodeMap[to_string_panic])
		}
	}()

	if c.err != nil {
		return "", c.err
	}

	//
	if decimalOrCoin == nil {
		return "", errors.New(errCodeMap[decimalorcoin_required])
	}

	//
	decimal, err := getDecimal(decimalOrCoin)
	if err != nil {
		return "", err
	}

	amountStr = c.amount.FloatString(decimal)

	return amountStr, nil
}

// @title: Output json.Number
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) ToJsonNumber(decimalOrCoin *interface{}) (amountJsonNumber json.Number, err error) {
	amountStr, err := c.ToString(decimalOrCoin)
	if err != nil {
		return json.Number(0), c.err
	}

	amountJsonNumber = json.Number(amountStr)

	return amountJsonNumber, nil
}

// @title: Output *big.Rat
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) ToBigRat() (*big.Rat, error) {
	return c.amount, c.err
}

// @title: Output *big.Int
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func (c *AmountDecimal) ToBigInt(decimalOrCoin interface{}) (*big.Int, error) {
	decimal, err := getDecimal(decimalOrCoin)
	if err != nil {
		return nil, err
	}

	expStr := fmt.Sprintf("1e%d", decimal)
	bigexp := new(big.Rat)
	_, ok := bigexp.SetString(expStr)
	if !ok {
		return nil, errors.New(errCodeMap[bigrat_setstring_fail])
	}

	result := new(bigRat).Mul(c.amount, bigexp)
	amountBigInt, ok := new(big.Int).SetString(result.FloatString(0), 10)
	if !ok {
		return nil, errors.New(errCodeMap[bigint_setstring_fail])
	}

	return amountBigInt, nil
}
