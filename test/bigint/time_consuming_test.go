package biginttest

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

// amountdecimal与decimal耗时测试
func TestBigIntTimeConsuming(t *testing.T) {
	BigIntTimeConsuming1()
	fmt.Println("")
	BigIntTimeConsuming2()
}

// amountdecimal耗时
func BigIntTimeConsuming1() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randBigIntList(bigIntMin, bigIntMax, bigIntNub)
	for _, v := range list {
		v2 := randBigInt(0, bigIntMax)
		amountdecimal.New(v).Add(v2).ToString(bigIntDecimal)
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}

// decimal耗时
func BigIntTimeConsuming2() {
	start := time.Now().UnixNano()
	fmt.Println("开始时间：", start)

	list := randBigIntList(bigIntMin, bigIntMax, bigIntNub)
	for _, v := range list {
		v2 := randBigInt(0, bigIntMax)
		v2d := decimal.NewFromBigInt(v2, int32(bigIntDecimal))
		decimal.NewFromBigInt(v, int32(bigIntDecimal)).Add(v2d).StringFixed(int32(bigIntDecimal))
	}

	end := time.Now().UnixNano()
	fmt.Println("结束时间：", end)
	fmt.Println("耗时：", end-start)
}