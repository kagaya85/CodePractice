package main

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	candidate2, max2 := []int{nums[n-1]}, math.MinInt64

	for i := n - 1; i >= 0; i-- {
		if nums[i] < max2 {
			return true
		}
		for len(candidate2) > 0 && nums[i] > candidate2[len(candidate2)-1] {
			max2 = candidate2[len(candidate2)-1]
			candidate2 = candidate2[:len(candidate2)-1]
		}
		if nums[i] > max2 {
			candidate2 = append(candidate2, nums[i])
		}
	}
	return false
}
