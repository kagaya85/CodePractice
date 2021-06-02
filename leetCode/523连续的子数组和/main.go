package main

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	remainder := 0
	rmp := map[int]int{0: -1}
	for i, num := range nums {
		remainder = (remainder + num) % k
		if preidx, has := rmp[remainder]; has {
			if i-preidx > 1 {
				return true
			}
		} else {
			rmp[remainder] = i
		}
	}

	return false
}
