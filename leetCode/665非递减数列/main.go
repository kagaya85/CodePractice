package main

func checkPossibility(nums []int) bool {
	count, n := 0, len(nums)

	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i-1] > nums[i+1] {
				// 前一个元素也大于后一个元素，则必须调整后一个元素
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
