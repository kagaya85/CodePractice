package main

func numWays(steps, arrLen int) int {
	const mod = 1e9 + 7
	n := min(arrLen, steps/2+1)
	dp := make([]int, n)

	dp[0] = 1

	for i := 1; i <= steps; i++ {
		next := make([]int, n)
		for j := range dp {
			next[j] = dp[j]
			if j > 0 {
				next[j] = (next[j] + dp[j-1]) % mod
			}
			if j < n-1 {
				next[j] = (next[j] + dp[j+1]) % mod
			}
		}
		dp = next
	}

	return dp[0]
}

func min(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
