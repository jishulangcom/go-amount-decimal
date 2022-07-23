package amountdecimal

import (
	"fmt"
	"math/big"
)

// @title: 获取*big.Int的实际金额
// @param: amount *big.Int
// @param: decimal int
// @return: string, error
// @auth: 技术狼(jishulang.com)
// @date: 2022/7/21 21:58
func BigIntActualAmount(amount *big.Int, decimal int) (string, error) {
	amountBigRat, err := amountRat(amount)
	if err != nil {
		return "", err
	}

	expStr := fmt.Sprintf("1e-%d", decimal)
	decimalBigRat, err := newBigRat(expStr)
	if err != nil {
		return "", err
	}

	result := new(bigRat).Mul(amountBigRat, decimalBigRat)
	return result.FloatString(decimal), nil
}
