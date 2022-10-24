func lengthOfLongestSubstring(s string) int {
	ans := math.MinInt
	count := map[byte]int{}
	l := 0

	for r, ch := range []byte(s) {
		count[ch]++
		for count[ch] > 1 {
			count[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	if ans == math.MinInt {
		return 0
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}