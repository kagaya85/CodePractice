package main

func maxTurbulenceSize(arr []int) (ans int) {
	n := len(arr)

	if n < 2 {
		return n
	}

	dp := make([][2]int, n)
	dp[0] = [2]int{1, 1}
	for i := 1; i < n; i++ {
		dp[i] = [2]int{1, 1}
		if arr[i-1] > arr[i] {
			dp[i][0] = dp[i-1][1] + 1
			ans = max(ans, dp[i][0])
		} else if arr[i-1] < arr[i] {
			dp[i][1] = dp[i-1][0] + 1
			ans = max(ans, dp[i][1])
		}
	}
	ans = max(ans, 1)
	return
}

func max(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
