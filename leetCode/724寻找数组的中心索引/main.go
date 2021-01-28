package main

func pivotIndex(nums []int) int {
	var right int
	for _, v := range nums {
		right += v
	}

	var left int
	for i, v := range nums {
		right -= v
		if left == right {
			return i
		}
		left += v
	}
	return -1
}
