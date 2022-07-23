package amountdecimal

import (
	"errors"
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

// @title: 输出*big.Int
// @return: string, error
// @description:
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func (c *AmountDecimal) BigInt() (*big.Int, error) {
	res, ok := new(big.Int).SetString(c.amount.FloatString(0), c.decimal)
	if !ok {
		return nil, errors.New(errCodeMap[bigint_setstring_fail])
	}
	return res, nil

}
