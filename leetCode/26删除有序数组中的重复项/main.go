package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	var p, q int
	for p, q = 0, 1; q < len(nums); q++ {
		if nums[p] != nums[q] {
			p++
			nums[p] = nums[q]
		}
	}

	return p + 1
}
