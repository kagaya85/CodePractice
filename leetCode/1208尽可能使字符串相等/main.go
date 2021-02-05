package main

func equalSubstring(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	diff := make([]int, n)
	for i := range s {
		diff[i] = abs(int(s[i]) - int(t[i]))
	}

	sum, start := 0, 0
	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
