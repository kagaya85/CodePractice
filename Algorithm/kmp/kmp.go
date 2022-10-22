package kmp

func Search(source, target string) int {
	s, t := 0, 0
	next := getNext(target)

	for s < len(source) {
		if source[s] == target[t] {
			s++
			t++
		} else if t != 0 {
			t = next[t-1]
		} else {
			s++
		}

		if t == len(target) {
			return s - t
		}
	}

	return -1
}

func getNext(s string) (next []int) {
	n := len(s)
	next = append(next, 0)
	idx, cur := 1, 0
	for idx < n {
		if s[cur] == s[idx] {
			cur++
			idx++
			next = append(next, cur)
		} else if cur != 0 {
			cur = next[cur]
		} else {
			idx++
			next = append(next, 0)
		}
	}
	return
}
