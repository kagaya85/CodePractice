package main

import "sort"

func arrayPairSum(nums []int) (sum int) {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i += 2 {
		sum += nums[i]
	}
	return
}
