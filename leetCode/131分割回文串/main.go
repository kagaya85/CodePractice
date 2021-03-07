package main

func partition(s string) (ans [][]string) {
	n := len(s)

	// -1 不是回文，0 未搜索, 1 是回文
	dp := make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, n)
	}

	var isPalindrome func(int, int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}

		// dp[i][j] == 0
		if s[i] == s[j] {
			dp[i][j] = isPalindrome(i+1, j-1)
		} else {
			dp[i][j] = -1
		}

		return dp[i][j]
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			ans = append(ans, append([]string{}, splits...))
		}
		for j := idx; j < n; j++ {
			if isPalindrome(idx, j) == 1 {
				splits = append(splits, s[idx:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)

	return
}
