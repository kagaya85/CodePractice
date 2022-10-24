func findMaxLength(nums []int) int {
	maxLength := math.MinInt
	sum := map[int]int{0: -1}
	pre := 0

	for i, v := range nums {
		if v == 0 {
			pre--
		} else {
			pre++
		}
		if idx, ok := sum[pre]; ok {
			maxLength = max(maxLength, i-idx)
		} else {
			sum[pre] = i
		}
	}

	if maxLength == math.MinInt {
		return 0
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}