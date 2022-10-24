func checkInclusion(s1 string, s2 string) bool {
	count := make([]int, 26)
	l, r := 0, 0
	for _, c := range s1 {
		count[c-'a']--
	}

	for r < len(s2) {
		idx := s2[r] - 'a'
		count[idx]++
		for count[idx] > 0 {
			count[s2[l]-'a']--
			l++
		}
		if r-l+1 == len(s1) {
			return true
		}
		r++
	}
	return false
}