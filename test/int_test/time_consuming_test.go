package inttest

import (
	"fmt"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestIntTimeConsuming(t *testing.T) {
	intTimeConsuming1()
	//fmt.Println("")
	//intTimeConsuming2()
	fmt.Println("")
	intTimeConsuming3()
}

// amountdecimal耗时
func intTimeConsuming1() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randIntList(intMin, intMax, intNub)
	for _, v := range list {
		v2 := randInt(intMin, intMax)
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
func intTimeConsuming2() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randIntList(intMin, intMax, intNub)
	for _, v := range list {
		v2 := randInt(0, intMin)
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
func intTimeConsuming3() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randIntList(intMin, intMax, intNub)
	for _, v := range list {
		v2 := randInt(0, intMin)
		calculation_decimal(add, v, v2)
		calculation_decimal(sub, v, v2)
		calculation_decimal(mul, v, v2)
		calculation_decimal(div, v, v2)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}