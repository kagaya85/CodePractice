package main

import (
	"fmt"
	"math"
)

const maxInt = math.MaxInt32

func main() {
	leaves := "rrryyyrryyyrr"
	fmt.Println(minimumOperations(leaves))
}

func minimumOperations(leaves string) (ans int) {

	n := len(leaves)
	f := make([][3]int, n)

	f[0][0] = boolToInt(leaves[0] == 'y')
	f[0][1] = maxInt
	f[0][2] = maxInt
	f[1][2] = maxInt

	for i := 1; i < n; i++ {

		isRed := boolToInt(leaves[i] == 'r')
		isYellow := boolToInt(leaves[i] == 'y')
		f[i][0] = f[i-1][0] + isYellow
		f[i][1] = min(f[i-1][0], f[i-1][1]) + isRed
		if i >= 2 {
			f[i][2] = min(f[i-1][1], f[i-1][2]) + isYellow
		}
	}
	return f[n-1][2]

}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
