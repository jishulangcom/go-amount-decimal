package test

import (
	"fmt"
	"github.com/jishulangcom/go-amount-decimal"
	"testing"
)

func Test(t *testing.T) {
	amount := amountdecimal.New("-.0115987", 8)
	result, err := amount.Div("2").String()
	fmt.Println(result, err)
}
