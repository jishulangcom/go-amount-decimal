package inttest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/shopspring/decimal"
	"math"
	"testing"
)

var (
	intMin     int = 0
	intMax     int = math.MaxInt64
	intNub     int = 10000
	intDecimal int = 20
)

func TestInt(t *testing.T) {
	list := randIntList(intMin, intMin, intDecimal)
	for _, v := range list {
		v2 := randInt(intMin, intMin)
		IntAdd(v, v2)
		IntSub(v, v2)
		IntMul(v, v2)
		IntDiv(v, v2)
	}
}

func IntAdd(v int, v2 int) {
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

func IntSub(v int, v2 int) {
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

func IntMul(v int, v2 int) {
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

func IntDiv(v int, v2 int) {
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

func calculation_amountdecimal(f uint8, v int, v2 int) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case add:
		return amountdecimal.New(v).Add(v2).ToString(intDecimal)
	case sub:
		return amountdecimal.New(v).Sub(v2).ToString(intDecimal)
	case mul:
		return amountdecimal.New(v).Mul(v2).ToString(intDecimal)
	case div:
		return amountdecimal.New(v).Div(v2).ToString(intDecimal)
	}

	return "", nil
}

func calculation_amountdecimal2(f uint8, v int, v2 int) (string, error) {
	// 计算方法所对应的操作
	switch f {
	case add:
		return amountdecimal.NewInt(v).AddInt(v2).ToString(intDecimal)
	case sub:
		return amountdecimal.NewInt(v).SubInt(v2).ToString(intDecimal)
	case mul:
		return amountdecimal.NewInt(v).MulInt(v2).ToString(intDecimal)
	case div:
		return amountdecimal.NewInt(v).DivInt(v2).ToString(intDecimal)
	}

	return "", nil
}

func calculation_decimal(f uint8, v int, v2 int) (string, error) {
	var res string

	v2d := decimal.NewFromInt(int64(v2))

	// 计算方法所对应的操作
	switch f {
	case add:
		res = decimal.NewFromInt(int64(v)).Add(v2d).StringFixed(int32(intDecimal))
	case sub:
		res = decimal.NewFromInt(int64(v)).Sub(v2d).StringFixed(int32(intDecimal))
	case mul:
		res = decimal.NewFromInt(int64(v)).Mul(v2d).StringFixed(int32(intDecimal))
	case div:
		res = decimal.NewFromInt(int64(v)).Div(v2d).StringFixed(int32(intDecimal))
	}

	return res, nil
}