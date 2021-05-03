package main

import "math"

func reverse(x int) (ans int) {
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		n := x % 10
		x /= 10
		ans = ans*10 + n
	}
	return
}
