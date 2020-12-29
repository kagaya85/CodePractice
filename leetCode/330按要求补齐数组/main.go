package main

func minPatches(nums []int, n int) (ans int) {
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			ans++
			x *= 2
		}
	}
	return
}
