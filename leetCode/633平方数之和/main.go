package main

import "math"

func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		b := math.Sqrt(float64(c - a*a))
		if b == math.Floor(b) {
			return true
		}
	}
	return false
}
