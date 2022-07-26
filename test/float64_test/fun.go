package float64test

import "math/rand"

func randFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randFloat64List(min, max float64, n int) []float64 {
	list := make([]float64, n)
	for i := range list {
		list[i] = randFloat64(min, max)
	}
	return list
}