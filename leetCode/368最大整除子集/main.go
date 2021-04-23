package main

import "sort"

func largestDivisibleSubset(nums []int) (ans []int) {
	n := len(nums)
	sort.Ints(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1
	}

	maxSize, maxNum := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] > maxSize {
			maxSize, maxNum = dp[i], nums[i]
		}
	}

	if maxSize == 1 {
		return []int{nums[0]}
	}

	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxNum%nums[i] == 0 {
			ans = append(ans, nums[i])
			maxNum = nums[i]
			maxSize--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
