package main

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	for l, r, tot := 0, 1, 0; r < len(nums); r++ {
		tot += (nums[r] - nums[r-1]) * (r - l)
		for tot > k {
			tot -= nums[r] - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
