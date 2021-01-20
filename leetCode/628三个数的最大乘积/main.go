package main

import "sort"

func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return max(nums[n-3]*nums[n-2]*nums[n-1], nums[0]*nums[1]*nums[n-1])
}

func max(arr ...int) (res int) {
	res = arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return
}
