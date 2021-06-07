package main

func findTargetSumWays(nums []int, target int) (count int) {
	n := len(nums)
	var dfs func(int, int)
	dfs = func(idx, sum int) {
		if idx >= n {
			if sum == target {
				count++
			}
			return
		}
		dfs(idx+1, sum+nums[idx])
		dfs(idx+1, sum-nums[idx])
	}
	dfs(0, 0)
	return
}
