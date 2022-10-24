func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k == 0 || k == 1 {
		return 0
	}
	n := len(nums)
	l, r := 0, 0
	mul := 1
	for r < n {
		mul *= nums[r]
		for mul >= k {
			mul /= nums[l]
			l++
		}
		ans += r - l + 1
		r++
	}
	return
}