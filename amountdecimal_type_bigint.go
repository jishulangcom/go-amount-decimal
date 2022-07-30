package amountdecimal

import (
	"errors"
	"fmt"
	"math/big"
)

// @title: initialization *big.Int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func NewBigInt(amount *big.Int) *AmountDecimal {
	return NewString(amount.String())
}

// @title: addition *big.Int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) AddBigInt(amounts ...*big.Int) *AmountDecimal {
	return c.amountsBigInt(add, amounts...)
}

// @title: subtraction *big.Int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) SubBigInt(amounts ...*big.Int) *AmountDecimal {
	return c.amountsBigInt(sub, amounts...)
}

// @title: multiplication *big.Int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) MulBigInt(amounts ...*big.Int) *AmountDecimal {
	return c.amountsBigInt(mul, amounts...)
}

// @title: division *big.Int type
// @auth: jishulang.com
// @date: 2022/7/30 22:52
func (c *AmountDecimal) DivBigInt(amounts ...*big.Int) *AmountDecimal {
	return c.amountsBigInt(div, amounts...)
}

func (c *AmountDecimal) amountsBigInt(f uint8, amounts ...*big.Int) *AmountDecimal {
	if c.err != nil {
		return c
	}

	var ad *AmountDecimal
	ad = c
	for _, amount := range amounts {
		if f == div && ad.amount == bigrat_zero {
			return ad
		}

		if f == div && amount == bigint_zero {
			ad.err = errors.New(errCodeMap[amount_divisor_cannot])
			return ad
		}

		amountDecimal := NewBigInt(amount)
		if amountDecimal.err != nil {
			return amountDecimal
		}

		ad = bigRatCalculation(f, ad.amount, amountDecimal.amount)
		if ad.err != nil {
			return ad
		}
	}

	return ad
}

//----------------------------------------------------------------------------------------------------------------------
// @title: Get the actual amount of *big.int
// @auth: jishulang.com
// @date: 2022/7/21 21:58
func BigIntActualAmount(amount *big.Int, decimalOrCoin interface{}) (string, error) {
	decimal, err := getDecimal(decimalOrCoin)
	if err != nil {
		return "", err
	}

	ad := NewBigInt(amount)
	if ad.err != nil {
		return "", ad.err
	}

	expStr := fmt.Sprintf("1e-%d", decimal)
	expAd := NewString(expStr)
	if expAd.err != nil {
		return "", expAd.err
	}

	result := new(bigRat).Mul(ad.amount, expAd.amount)
	return result.FloatString(decimal), nil
}
