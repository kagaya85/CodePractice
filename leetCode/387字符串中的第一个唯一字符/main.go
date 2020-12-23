package main

func firstUniqChar(s string) int {
	charCnt := [26]int{}
	for _, c := range s {
		charCnt[c-'a']++
	}
	for i, c := range s {
		if charCnt[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
