package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			leftdown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftdown = leftdown + 1
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, leftdown)
		}
	}

	return dp[m][n]
}

func min(v ...int) int {
	min := v[0]
	for _, n := range v[1:] {
		if n < min {
			min = n
		}
	}
	return min
}
