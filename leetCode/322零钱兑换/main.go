func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			tmp := i - c
			if tmp >= 0 && dp[tmp] != -1 {
				dp[i] = min(dp[i], dp[tmp]+1)
			}
		}
		if dp[i] == math.MaxInt {
			dp[i] = -1
		}
	}

	return dp[amount]
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