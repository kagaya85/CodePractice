package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	// 外层遍历硬币数量
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
