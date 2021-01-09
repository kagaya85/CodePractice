package main

func maxProfit(prices []int) int {
	n := len(prices)
	dp := [4]int{-prices[0], 0, -prices[0], 0}

	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], -prices[i])
		dp[1] = max(dp[1], dp[0]+prices[i])
		dp[2] = max(dp[2], dp[1]-prices[i])
		dp[3] = max(dp[3], dp[2]+prices[i])
	}

	return dp[3]
}

func max(arr ...int) (ans int) {
	ans = arr[0]
	for _, v := range arr[1:] {
		if v > ans {
			ans = v
		}
	}
	return
}
