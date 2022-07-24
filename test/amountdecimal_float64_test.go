package test

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/shopspring/decimal"
	"math/rand"
	"testing"
	"time"
)

var (
	float64Min float64 = 0
	float64Max float64 = 99999999999999999999999999999999999999999999999999.999999999999999999999999999999999999999999
	float64N int = 10000
	decimalInt int = 22
)


func TestFloat64(t *testing.T) {
	float64List := randFloat64List(float64Min, float64Max, float64N)
	for _, v := range float64List {
		v2 := RandFloat64(0, float64Max)
		fmt.Printf("\n\r ----------%v + %v ------------------------\n", v, v2)


		res, err := amountdecimal.New(v).Add(v2).ToString(decimalInt)
		fmt.Printf("AmountDecimal result:%v, %v \n", err, res)

		res2, err2 := amountdecimal.New(v).AddFloat64(v2).ToString(decimalInt)
		fmt.Printf("AmountDecimal result2:%v, %v \n", err2, res2)

		v3d := decimal.NewFromFloat(v2)
		res3 := decimal.NewFromFloat(v).Add(v3d).StringFixed(int32(decimalInt))
		fmt.Printf("Decimal result:%v, %v \n", nil, res3)

		if res != res2 {
			//panic("AmountDecimal NewFloat64与New不相等")
		}

		if res != res3 {
			panic("AmountDecimal与Decimal不相等")
		}
	}
}

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

	float64List := randFloat64List(float64Min, float64Max, float64N)
	for _, v := range float64List {
		v2 := RandFloat64(0, float64Max)
		amountdecimal.New(v).Add(v2).ToString(decimalInt)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func float64TimeConsuming2() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	float64List := randFloat64List(float64Min, float64Max, float64N)
	for _, v := range float64List {
		v2 := RandFloat64(0, float64Max)
		v2d := decimal.NewFromFloat(v2)
		decimal.NewFromFloat(v).Add(v2d).StringFixed(int32(decimalInt))
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
	return min + rand.Float64() * (max - min)
}

func randFloat32List(min, max float32, n int) []float32 {
	list := make([]float32, n)
	for i := range list {
		list[i] = min + rand.Float32() * (max - min)
	}
	return list
}
