package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums...)
	}
	_rob := func(nums []int) int {
		p, q := nums[0], max(nums[0], nums[1])
		for _, n := range nums[2:] {
			p, q = q, max(p+n, q)
		}
		return q
	}
	return max(_rob(nums[1:]), _rob(nums[:n-1]))
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
