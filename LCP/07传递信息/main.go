package main

func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 1; i <= k; i++ {
		for _, r := range relation {
			from, to := r[0], r[1]
			dp[i][to] += dp[i-1][from]
		}
	}
	return dp[k][n-1]
}
