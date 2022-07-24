package biginttest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

var (
	bigIntMin int64 = 0
	bigIntMax int64 = 999999
	bigIntNub int = 10000
	bigIntDecimal int = 22
)


func TestBigInt(t *testing.T) {
	list := randBigIntList(bigIntMin, bigIntMax, bigIntNub)
	for _, v := range list {
		add(v)
		sub(v)
		mul(v)
		div(v)
	}
}

func add(v *big.Int)  {
	v2 := RandBigInt(0, bigIntMax)
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)


	res, err := amountdecimal.New(v).Add(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal  result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.NewBigInt(v).AddBigInt(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)

	v3d := decimal.NewFromBigInt(v2, int32(bigIntDecimal))
	res3 := decimal.NewFromBigInt(v, int32(bigIntDecimal)).Add(v3d).StringFixed(int32(bigIntDecimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func sub(v *big.Int)  {
	v2 := RandBigInt(0, bigIntMax)
	fmt.Printf("\n\r ----------%v - %v ------------------------\n", v, v2)


	res, err := amountdecimal.New(v).Sub(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.NewBigInt(v).SubBigInt(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)

	v3d := decimal.NewFromBigInt(v2, int32(bigIntDecimal))
	res3 := decimal.NewFromBigInt(v, int32(bigIntDecimal)).Sub(v3d).StringFixed(int32(bigIntDecimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func mul(v *big.Int)  {
	v2 := RandBigInt(0, bigIntMax)
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)


	res, err := amountdecimal.New(v).Mul(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.NewBigInt(v).MulBigInt(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)

	v3d := decimal.NewFromBigInt(v2, int32(bigIntDecimal))
	res3 := decimal.NewFromBigInt(v, int32(bigIntDecimal)).Mul(v3d).StringFixed(int32(bigIntDecimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func div(v *big.Int)  {
	v2 := RandBigInt(0, bigIntMax)
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)


	res, err := amountdecimal.New(v).Div(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.NewBigInt(v).DivBigInt(v2).ToString(bigIntDecimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)

	v3d := decimal.NewFromBigInt(v2, int32(bigIntDecimal))
	res3 := decimal.NewFromBigInt(v, int32(bigIntDecimal)).Div(v3d).StringFixed(int32(bigIntDecimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal New与NewBigInt不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}