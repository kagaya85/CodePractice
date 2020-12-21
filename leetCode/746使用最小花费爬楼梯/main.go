package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(arr ...int) (res int) {
	res = arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
