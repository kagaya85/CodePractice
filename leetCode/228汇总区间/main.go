package main

import (
	"sort"
	"strconv"
)

func summaryRanges(nums []int) (ans []string) {
	// 有序数组
	sort.Ints(nums)

	for i, n := 0, len(nums); i < n; {
		start := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		span := strconv.Itoa(nums[start])
		if start < i-1 {
			span += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, span)
	}
	return
}
