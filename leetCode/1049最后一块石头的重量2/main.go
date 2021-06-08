package main

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	sum := 0
	for _, w := range stones {
		sum += w
	}
	m := sum / 2
	dp := make([][]bool, n+1) // 对于dp[i][j] 前i个石头是否可以凑出重量j
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i, w := range stones {
		for j := 0; j <= m; j++ {
			if w > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j-w] || dp[i][j]
			}
		}
	}

	for j := m; j >= 0; j-- {
		if dp[n][j] {
			return sum - 2*j
		}
	}

	return sum
}
