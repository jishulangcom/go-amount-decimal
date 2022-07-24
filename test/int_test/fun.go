package inttest

import (
	"math/rand"
)

func randIntList(min, max int, n int) []int {
	list := make([]int, n)
	for i := range list {
		list[i] = randInt(min, max)
	}
	return list
}

func randInt(min, max int) int {
	return min + rand.Int()*(max-min)
}