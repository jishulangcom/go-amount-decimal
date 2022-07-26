package float64test

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/shopspring/decimal"
	"testing"
)

var (
	float64Min     float64 = 0
	float64Max     float64 = 99.9999999999999999999999999999999
	float64Nub     int     = 10000
	float64Decimal int     = 15
)


func TestInt(t *testing.T) {
	list := randFloat64List(float64Min, float64Max, float64Nub)
	for _, v := range list {
		v2 := randFloat64(float64Min, float64Max)
		Float64Add(v, v2)
		Float64Sub(v, v2)
		Float64Mul(v, v2)
		Float64Div(v, v2)
	}
}

func Float64Add(v float64, v2 float64) {
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)

	res, err := calculation_amountdecimal(add, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := calculation_amountdecimal2(add, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := calculation_decimal(add, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float64Sub(v float64, v2 float64) {
	fmt.Printf("\n\r ----------%v - %v ------------------------\n", v, v2)

	res, err := calculation_amountdecimal(sub, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := calculation_amountdecimal2(sub, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := calculation_decimal(sub, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float64Mul(v float64, v2 float64) {
	fmt.Printf("\n\r ----------%v * %v ------------------------\n", v, v2)

	res, err := calculation_amountdecimal(mul, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := calculation_amountdecimal2(mul, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		panic(err2)
	}

	res3, err3 := calculation_decimal(mul, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func Float64Div(v float64, v2 float64) {
	fmt.Printf("\n\r ----------%v / %v ------------------------\n", v, v2)

	res, err := calculation_amountdecimal(div, v, v2)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		fmt.Println("calculation_amountdecimal err:", err)
		panic(err)
	}

	res2, err2 := calculation_amountdecimal2(div, v, v2)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err2 != nil {
		fmt.Println("calculation_amountdecimal2 err:", err2)
		panic(err2)
	}

	res3, err3 := calculation_decimal(div, v, v2)
	fmt.Printf("Decimal result:%v, %v \n", err3, res3)

	if res != res2 {
		panic("AmountDecimal New与NewInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func calculation_amountdecimal(f uint8, v float64, v2 float64) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case add:
		return amountdecimal.New(v).Add(v2).ToString(float64Decimal)
	case sub:
		return amountdecimal.New(v).Sub(v2).ToString(float64Decimal)
	case mul:
		return amountdecimal.New(v).Mul(v2).ToString(float64Decimal)
	case div:
		return amountdecimal.New(v).Div(v2).ToString(float64Decimal)
	}

	return "", nil
}

func calculation_amountdecimal2(f uint8, v float64, v2 float64) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case add:
		return amountdecimal.NewFloat64(v).AddFloat64(v2).ToString(float64Decimal)
	case sub:
		return amountdecimal.NewFloat64(v).SubFloat64(v2).ToString(float64Decimal)
	case mul:
		return amountdecimal.NewFloat64(v).MulFloat64(v2).ToString(float64Decimal)
	case div:
		return amountdecimal.NewFloat64(v).DivFloat64(v2).ToString(float64Decimal)
	}

	return "", nil
}

func calculation_decimal(f uint8, v float64, v2 float64) (string, error) {
	var res string

	v2d := decimal.NewFromFloat(v2)

	// 计算方法所对应的操作
	switch f {
	case add:
		res = decimal.NewFromFloat(v).Add(v2d).StringFixed(int32(float64Decimal))
	case sub:
		res = decimal.NewFromFloat(v).Sub(v2d).StringFixed(int32(float64Decimal))
	case mul:
		res = decimal.NewFromFloat(v).Mul(v2d).StringFixed(int32(float64Decimal))
	case div:
		res = decimal.NewFromFloat(v).Div(v2d).StringFixed(int32(float64Decimal))
	}

	return res, nil
}