package main

import (
	"math"
	"math/bits"
)

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	m := 1 << n
	sum := make([]int, m)
	for i := 1; i < m; i++ {
		x := bits.TrailingZeros(uint(i))
		y := i ^ 1<<x
		sum[i] = sum[y] + jobs[x]
	}

	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i, s := range sum {
		dp[0][i] = s
	}
	for i := 1; i < k; i++ {
		for j := 0; j < m; j++ {
			minTime := math.MaxInt64
			// 当前i个工人分配情况为j时，遍历j的所有子集
			for x := j; x > 0; x = (x - 1) & j {
				// 第i个人分配x的工作，前i-1个分配j-x
				minTime = min(minTime, max(dp[i-1][j-x], sum[x]))
			}
			dp[i][j] = minTime
		}
	}
	return dp[k-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
