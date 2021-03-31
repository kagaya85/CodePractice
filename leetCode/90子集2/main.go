package main

import "sort"

func subsetsWithDup(nums []int) (ans [][]int) {

	l := len(nums)
	sort.Ints(nums)
	sub := []int{}

	var dfs func(pre bool, cur int)
	dfs = func(pre bool, cur int) {
		if cur == l {
			ans = append(ans, append([]int(nil), sub...))
			return
		}
		dfs(false, cur+1)
		// 如果没选择前一个数字且当前数字与前一个相同, 跳过
		if !pre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		sub = append(sub, nums[cur])
		dfs(true, cur+1)
		sub = sub[:len(sub)-1]
	}

	dfs(false, 0)

	return
}
