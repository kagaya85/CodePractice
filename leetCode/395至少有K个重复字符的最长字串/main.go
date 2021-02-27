package main

import "strings"

func longestSubstring(s string, k int) (ans int) {
	if len(s) < k {
		return
	}

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}

	var split byte
	for i, num := range cnt {
		if 0 < num && num < k {
			split = 'a' + byte(i)
		}
	}

	if split == 0 {
		return len(s)
	}

	for _, sub := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(sub, k))
	}

	return
}

func max(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
