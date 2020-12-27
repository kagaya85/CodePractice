package main

func isIsomorphic(s, t string) bool {
	sl, tl := len(s), len(t)
	if sl != tl {
		return false
	}

	m := make(map[byte]byte)
	n := make(map[byte]byte)

	for i := range s {
		if c, has := m[s[i]]; has {
			if c != t[i] {
				return false
			}
		}
		if c, has := n[t[i]]; has {
			if c != s[i] {
				return false
			}
		}
		m[s[i]] = t[i]
		n[t[i]] = s[i]
	}
	return true
}
