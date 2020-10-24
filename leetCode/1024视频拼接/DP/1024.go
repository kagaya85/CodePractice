package main

import "math"

func main() {

}

func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	dp[0] = 0
	for i := 1; i <= T; i++ {
		dp[i] = math.MaxInt32
	}

	for i := range dp {
		for _, clip := range clips {
			if i > clip[0] && i <= clip[1] && dp[clip[0]]+1 < dp[i] {
				dp[i] = dp[clip[0]] + 1
			}
		}
	}

	if dp[T] == math.MaxInt32 {
		return -1
	}

	return dp[T]
}
