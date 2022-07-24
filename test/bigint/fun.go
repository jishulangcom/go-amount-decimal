package biginttest

import (
	"math/big"
	"math/rand"
)

func randBigIntList(min, max int64, n int) []*big.Int {
	list := make([]*big.Int, n)
	for i := range list {
		list[i] = RandBigInt(min, max)
	}
	return list
}

func RandBigInt(min, max int64) *big.Int {
	i := min + int64(rand.Int()) * (max - min)
	return new(big.Int).SetInt64(i)
}

