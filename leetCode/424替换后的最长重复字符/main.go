package main

func characterReplacement(s string, k int) int {
	maxSame, left := 0, 0
	count := make([]int, 26)

	for right, ch := range s {
		idx := ch - 'A'
		count[idx]++
		if count[idx] > maxSame {
			maxSame = count[idx]
		}
		if right-left+1-maxSame > k {
			count[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}
