package main

import "math"

func minCut(s string) int {
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

	cnt := make([]int, n)
	for i := range cnt {
		if isPalindrome(0, i) == 1 {
			continue
		}
		cnt[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if isPalindrome(j+1, i) == 1 && cnt[j]+1 < cnt[i] {
				cnt[i] = cnt[j] + 1
			}
		}
	}
	return cnt[n-1]
}
