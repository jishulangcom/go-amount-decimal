package test

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal"
	"testing"
)

func Test(t *testing.T) {
	amount := amountdecimal.New("0.115987", 8)
	result, err := amount.Add("1.397").Sub("1").String()
	fmt.Println(result, err)
}
