package main

import "fmt"

func main() {
	S := "ababcbacadefegdehijhklij"

	fmt.Println(partitionLabels(S))
}

func partitionLabels(S string) (ans []int) {
	lastPos := make([]int, 26)

	for i, c := range S {
		lastPos[c-'a'] = i
	}

	start, end := 0, 0

	for i, c := range S {
		end = max(end, lastPos[c-'a'])
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
