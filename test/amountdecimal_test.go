package test

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal"
	"math/big"
	"testing"
)

type kk struct {

}

func Test(t *testing.T) {
	//var i float32 = 1
	//i := json.Number("14")
	//i := &struct {}{}
	//var i interface{}
	amountBigRat := new(big.Rat)
	amountBigRat.SetString("1.222")

	amountBigInt := new(big.Int)
	amountBigInt.SetString("111", 0)

	//i := &kk{}


	amount := amountdecimal.New(amountBigInt, 8)
	result, err := amount.Div("2.1").String()
	fmt.Println(result, err)
}
