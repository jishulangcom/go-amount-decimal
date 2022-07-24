package float_test

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/shopspring/decimal"
	"math/rand"
	"testing"
	"time"
)

var (
	float64Min     float64 = 0
	float64Max     float64 = 99.9999999999999999999999999999999
	float64Nub     int     = 10000
	float64Decimal int     = 15
)

func TestFloat64(t *testing.T) {
	float64List := randFloat64List(float64Min, float64Max, float64Nub)
	for _, v := range float64List {
		float64Add(v)
		float64Sub(v)
		float64Mul(v)
		float64Div(v)
	}
}

func float64Add(v float64) {
	v2 := RandFloat64(0, float64Max)
	fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)

	res, err := amountdecimal.New(v).Add(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.New(v).AddFloat64(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err != nil {
		panic(err2)
	}

	v3d := decimal.NewFromFloat(v2)
	res3 := decimal.NewFromFloat(v).Add(v3d).StringFixed(int32(float64Decimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal NewFloat64与New不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func float64Sub(v float64) {
	v2 := RandFloat64(0, float64Max)
	fmt.Printf("\n\r ----------%v - %v ------------------------\n", v, v2)

	res, err := amountdecimal.New(v).Sub(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.New(v).SubFloat64(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err != nil {
		panic(err2)
	}

	v3d := decimal.NewFromFloat(v2)
	res3 := decimal.NewFromFloat(v).Sub(v3d).StringFixed(int32(float64Decimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal NewFloat64与New不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func float64Mul(v float64) {
	v2 := RandFloat64(0, float64Max)
	fmt.Printf("\n\r ----------%v * %v ------------------------\n", v, v2)

	res, err := amountdecimal.New(v).Mul(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.New(v).MulFloat64(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err != nil {
		panic(err2)
	}

	v3d := decimal.NewFromFloat(v2)
	res3 := decimal.NewFromFloat(v).Mul(v3d).StringFixed(int32(float64Decimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal NewFloat64与New不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}

func float64Div(v float64) {
	v2 := RandFloat64(0, float64Max)
	fmt.Printf("\n\r ----------%v / %v ------------------------\n", v, v2)

	res, err := amountdecimal.New(v).Div(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result:%v, %v \n", err, res)
	if err != nil {
		panic(err)
	}

	res2, err2 := amountdecimal.New(v).DivFloat64(v2).ToString(float64Decimal)
	fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)
	if err != nil {
		panic(err2)
	}

	v3d := decimal.NewFromFloat(v2)
	res3 := decimal.NewFromFloat(v).Div(v3d).StringFixed(int32(float64Decimal))
	fmt.Printf("Decimal result:%v, %v \n", nil, res3)

	if res != res2 {
		panic("AmountDecimal NewFloat64与New不相等")
	}

	if res != res3 {
		panic("AmountDecimal与Decimal不相等")
	}
}


//----------------------------------------------------------------------------------------------------------------------
// amountdecimal与decimal耗时测试
func TestFloat64TimeConsuming(t *testing.T) {
	float64TimeConsuming1()
	fmt.Println("")
	float64TimeConsuming2()
}

// amountdecimal耗时
func float64TimeConsuming1() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randFloat64List(float64Min, float64Max, float64Nub)
	for _, v := range list {
		v2 := RandFloat64(0, float64Max)
		amountdecimal.New(v).Add(v2).ToString(float64Decimal)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func float64TimeConsuming2() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randFloat64List(float64Min, float64Max, float64Nub)
	for _, v := range list {
		v2 := RandFloat64(0, float64Max)
		v2d := decimal.NewFromFloat(v2)
		decimal.NewFromFloat(v).Add(v2d).StringFixed(int32(float64Decimal))
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

func randFloat64List(min, max float64, n int) []float64 {
	list := make([]float64, n)
	for i := range list {
		list[i] = RandFloat64(min, max)
	}
	return list
}

func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randFloat32List(min, max float32, n int) []float32 {
	list := make([]float32, n)
	for i := range list {
		list[i] = min + rand.Float32()*(max-min)
	}
	return list
}
