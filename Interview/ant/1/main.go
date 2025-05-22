package main

import "strings"

// "42"
// "   -42"
// "00000042"
// "4193 with words"
func convertToNum(s string) int32 {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	var sign int32 = 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	num := int32(0)
	for _, c := range s {
		if c < '0' || c > '9' {
			break
		}
		num = num*10 + int32(c-'0')
	}

	return sign * num
}
