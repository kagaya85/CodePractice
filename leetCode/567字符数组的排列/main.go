package main

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	diff := 0
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, c := range s1 {
		cnt[c-'a']++
		cnt[s2[i]-'a']--
	}
	for _, num := range cnt {
		if num != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}

	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		cnt[x]--
		cnt[y]++
		if cnt[x] == 0 {
			diff--
		} else if cnt[x] == -1 {
			diff++
		}
		if cnt[y] == 0 {
			diff--
		} else if cnt[y] == 1 {
			diff++
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
