package float64test

import (
	"fmt"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestIntTimeConsuming(t *testing.T) {
	list := randFloat64List(float64Min, float64Max, float64Decimal)

	float64TimeConsuming1(list)
	fmt.Println("")
	float64TimeConsuming2(list)
	fmt.Println("")
	float64TimeConsuming3(list)
}

// amountdecimal耗时
func float64TimeConsuming1(list []float64) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := randFloat64(float64Min, float64Max)
		calculation_amountdecimal(add, v, v2)
		calculation_amountdecimal(sub, v, v2)
		calculation_amountdecimal(mul, v, v2)
		calculation_amountdecimal(div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// amountdecimal2耗时
func float64TimeConsuming2(list []float64) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := randFloat64(float64Min, float64Max)
		calculation_amountdecimal2(add, v, v2)
		calculation_amountdecimal2(sub, v, v2)
		calculation_amountdecimal2(mul, v, v2)
		calculation_amountdecimal2(div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func float64TimeConsuming3(list []float64) {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	for _, v := range list {
		v2 := randFloat64(float64Min, float64Max)
		calculation_decimal(add, v, v2)
		calculation_decimal(sub, v, v2)
		calculation_decimal(mul, v, v2)
		calculation_decimal(div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}