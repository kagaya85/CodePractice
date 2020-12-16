package main

import "strings"

func main() {

}

func wordPattern(pattern string, s string) bool {
	pm := make(map[byte]string)
	sm := make(map[string]byte)
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i, w := range words {
		s, has := pm[pattern[i]]
		if !has {
			if _, has := sm[w]; has {
				return false
			}
			pm[pattern[i]] = w
			sm[w] = pattern[i]
			continue
		}
		if s != w {
			return false
		}
	}
	return true
}
