package main

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
// 题目数据保证答案符合 32 位带符号整数范围。

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 0
	}

	// dp[i][j] s[:i]子序列中t[:j]出现的个数
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// t为空串的情况
	for i := range dp {
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
