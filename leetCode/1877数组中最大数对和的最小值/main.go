package main

import "sort"

func minPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for i, val := range nums[:n/2] {
		ans = max(ans, val+nums[n-i-1])
	}
	return
}

func max(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
