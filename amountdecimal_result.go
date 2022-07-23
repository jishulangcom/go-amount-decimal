package amountdecimal

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

// @title: 输出string
// @return: string, error
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) String() (amountStr string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(errCodeMap[string_fail])
		}
	}()

	if c.err != nil {
		return "", c.err
	}

	amountStr = c.amount.FloatString(c.decimal)
	return amountStr, nil
}

// @title: 输出json.Number
// @return: string, error
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) JsonNumber() (amountJsonNumber json.Number, err error) {
	amountStr, err := c.String()
	if err != nil {
		return json.Number(0), c.err
	}

	amountJsonNumber = json.Number(amountStr)

	return amountJsonNumber, nil
}

// @title: 输出*big.Rat
// @return: string, error
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) BigRat() (*big.Rat, error) {
	return c.amount, c.err
}

// @title: 输出*big.Int
// @return: string, error
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) BigInt() (*big.Int, error) {
	expStr := fmt.Sprintf("1e%d", c.decimal)
	bigexp := new(big.Rat)
	_, ok := bigexp.SetString(expStr)
	if !ok {
		return nil, errors.New("xxxxx")
	}

	result := new(bigRat).Mul(c.amount, bigexp)
	amountBigInt, ok :=new(big.Int).SetString(result.FloatString(0), 10)
	if !ok {
		return nil, errors.New("sssssssssssss")
	}

	return amountBigInt, nil
}