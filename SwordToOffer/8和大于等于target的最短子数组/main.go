func minSubArrayLen(target int, nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	left, right := 0, 0
	ans = math.MaxInt32
	for right < len(nums) {
		sum += nums[right]
		for sum >= target && left <= right {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return
}