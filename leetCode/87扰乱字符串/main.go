package main

func isScramble(s1, s2 string) bool {
	n := len(s1)
	// dp[i][j][len]在s1的第i个字符和s2的第j个字符开始，
	// 长度为len的两个字符串是否”和谐“
	// 0-未知 1-和谐 -1-不和谐
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
		}
	}

	var dfs func(int, int, int) int8
	dfs = func(i1, i2, len int) (res int8) {
		cur := &dp[i1][i2][len]
		if *cur != 0 {
			return *cur
		}

		defer func() {
			*cur = res
		}()

		x, y := s1[i1:i1+len], s2[i2:i2+len]
		// 直接判断是否相等
		if x == y {
			return 1
		}

		// 判断每个字符数量是否相等
		cnt := [26]int{}
		for i := range x {
			cnt[x[i]-'a']++
			cnt[y[i]-'a']--
		}
		for _, n := range cnt {
			if n != 0 {
				return -1
			}
		}

		// 枚举每一个分割位置
		for i := 1; i < len; i++ {
			// 不交换
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, len-i) == 1 {
				return 1
			}
			// 交换
			if dfs(i1, i2+len-i, i) == 1 && dfs(i1+i, i2, len-i) == 1 {
				return 1
			}
		}

		return -1
	}

	return dfs(0, 0, n) == 1
}
