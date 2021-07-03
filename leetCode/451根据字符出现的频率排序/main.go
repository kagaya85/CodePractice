package main

import "bytes"

func frequencySort(s string) string {
	count := map[byte]int{}
	maxFreq := 0
	for i := range s {
		count[s[i]]++
		maxFreq = max(maxFreq, count[s[i]])
	}
	bkt := make([][]byte, maxFreq+1)
	for ch, cnt := range count {
		bkt[cnt] = append(bkt[cnt], ch)
	}
	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range bkt[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
