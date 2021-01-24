package main

func findLengthOfLCIS(nums []int) (ans int) {
	var start int
	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		ans = func(i, j int) int {
			if i > j {
				return i
			}
			return j
		}(ans, i-start+1)
	}
	return
}
