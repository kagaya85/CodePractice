package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("abc"))
}

func permutation(s string) (ans []string) {
	n := len(s)
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	visited := make([]bool, n)
	tmp := make([]byte, 0, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, string(tmp))
			return
		}

		for i, c := range arr {
			if visited[i] || i > 0 && arr[i-1] == arr[i] && !visited[i-1] {
				continue
			}
			visited[i] = true
			tmp = append(tmp, c)
			backtrack(idx + 1)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}
	backtrack(0)
	return
}
