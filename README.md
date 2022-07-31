# Amount Decimal

Development based on golang 【math/big.rat】 package。

It is used to calculate the amount handling fee, exchange rate



## Features





## Install

Run `go get github.com/jishulangcom/go-amount-decimal`





## Requirements 

Decimal library requires Go version `>=1.7`





## Usage

```go
amount type：
    【"string", "json.Number"】
    【"float32", "float64"】
    【"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"】
    【"*big.Int", "*big.Rat"】

calculation func：
    .Add(amounts ...interface{}) // addition
    .Sub(amounts ...interface{}) // subtraction
    .Mul(amounts ...interface{}) // multiplication
    .Div(amounts ...interface{}) // division

result func: 
    .ToString(decimalOrCoin interface{})     // output string type
    .ToJsonNumber(decimalOrCoin interface{}) // output json.Number type
    .ToBigInt(decimalOrCoin interface{})     // output *big.Int type
    .ToBigRat()                              // output *big.Rat type
```

```go
package main

import (
	"encoding/json"
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"math/big"
)

func main() {
	var f float64 = 0.5 // 【"float32", "float64"】
	var s string = "1"
	var i int = 2 // 【"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"】
	var j json.Number = json.Number("3")
	bi := new(big.Int).SetInt64(4) // *big.Int
	var decimal int = 5

	//---------------------【Basic Usage】-----------------------------------------------------------------------
	// 0.5+1 = (1.5)-2 = (-0.5)*3 = (-1.5)/4
	// Calculated in order from left to right
	amountB1, errB1 := amountdecimal.New(f).Add(s).Sub(i).Mul(j).Div(bi).ToString(decimal)
	fmt.Println("【Basic Usage】：", amountB1, errB1) // -0.37500 <nil>

	//---------------------【Specify Amount Type】-----------------------------------------------------------
	// 0.5 +1 +2 +3 +4 +5
	amountS1, errS1 := amountdecimal.NewFloat64(f).AddInt(i).ToString(decimal)
	fmt.Println("【Specify Amount Type】：", amountS1, errS1) // 2.50000 <nil>

	//---------------------【Multiple Amounts】-----------------------------------------------------------
	// 0.5 +1 +2 +3 +4 +5
	amountM1, errM1 := amountdecimal.New(f).Add(s, i, j, bi).ToString(decimal)
	fmt.Println("【Multiple Amounts】：", amountM1, errM1) // 10.50000 <nil>

	//---------------------【Coin Decimal】------------------------------------------------------------------------
	amount := 100
	fee := 0.001
	amountC1, errC1 := amountdecimal.New(amount).Mul(fee).ToString("BTC")
	fmt.Println("【Coin Decimal1】：", amountC1, errC1) // 0.10000000 <nil>

	amountC2, errC2 := amountdecimal.New(amount).Mul(fee).ToBigInt(amountdecimal.COIN_BTC)
	fmt.Println("【Coin Decimal2】：", amountC2, errC2) // 10000000 <nil>

	amountC3, errC3 := amountdecimal.BigIntActualAmount(amountC2, amountdecimal.COIN_BTC_DECIMAL)
	fmt.Println("【Coin Decimal3】：", amountC3, errC3) // 10000000 <nil>
}
```





## Documentation

https://pkg.go.dev/github.com/jishulangcom/go-amount-decimal

https://github.com/jishulangcom/go-amount-decimal-test





## Production Usage

* If you are using this in production, please let us know!

  



## FAQ





## License

The MIT License (MIT)
