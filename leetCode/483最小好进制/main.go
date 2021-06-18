package main

import (
	"math"
	"math/bits"
	"strconv"
)

func smallestGoodBase(n string) string {
	val, _ := strconv.Atoi(n)
	max := bits.Len(uint(val)) - 1

	for m := max; m > 1; m-- {
		k := int(math.Pow(float64(val), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == val {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(val - 1)
}
