func largestRectangleArea(heights []int) (ans int) {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	stk := []int{}

	for l := 0; l < n; l++ {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[l] {
			stk = stk[:len(stk)-1]
		}
		left[l] = -1
		if len(stk) > 0 {
			left[l] = stk[len(stk)-1]
		}
		stk = append(stk, l)
	}

	stk = stk[:0]
	for r := n - 1; r >= 0; r-- {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[r] {
			stk = stk[:len(stk)-1]
		}
		right[r] = n
		if len(stk) > 0 {
			right[r] = stk[len(stk)-1]
		}
		stk = append(stk, r)
	}

	for i, h := range heights {
		ans = max(ans, (right[i]-left[i]-1)*h)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}