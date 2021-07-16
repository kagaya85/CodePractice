package main

import "sort"

func search(nums []int, target int) int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return 0
	}
	right := sort.SearchInts(nums, target+1)
	return right - left
}
