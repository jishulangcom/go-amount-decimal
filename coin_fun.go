package amountdecimal

import (
	"errors"
)

// @title: 获取精度
// @auth: jishulang.com
// @date: 2022/7/21 23:15
func getDecimal(decimalOrCoin interface{}) (int, error) {
	decimalType := getType(decimalOrCoin)

	if decimalType == type_string {
		coin := interfaceToStr(decimalOrCoin)
		coin = strToUpper(coin)
		if coinInfo, ok := CoinMap[coin]; ok {
			return coinInfo.Decimal, nil
		}

		return 0, errors.New(errCodeMap[coin_wrong])
	}

	if _, ok := numberTypeMap[decimalType]; ok {
		decimal, ok2 := decimalOrCoin.(int)
		if ok2 {
			return decimal, nil
		}
		return 0, errors.New(errCodeMap[coin_wrong])
	}

	return 0, errors.New(errCodeMap[coin_type_wrong])
}
