/*
	amountdecimal
	【package name:】Amount Decimal
	【auth:】jishulang.com
*/
package amountdecimal

import "math/big"

type AmountDecimal struct {
	amount *big.Rat
	err    error
}

// calculation func
const (
	add = iota // addition
	sub        // subtraction
	mul        // multiplication
	div        // division
)
