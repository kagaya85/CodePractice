func maxSubArray(nums []int) int {
	cur, res := 0, math.MinInt
	for _, v := range nums {
		if cur < 0 {
			cur = v
		} else {
			cur += v
		}
		res = max(res, cur)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}