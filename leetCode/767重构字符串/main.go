package main

import "fmt"

func main() {
	fmt.Println(reorganizeString("aab"))
}

func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	cnt := [26]int{}
	maxCnt := 0
	for _, v := range S {
		v -= 'a'
		cnt[v]++
		if cnt[v] > maxCnt {
			maxCnt = cnt[v]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	ans := make([]byte, n)
	evenIdx, oddIdx, half := 0, 1, n/2
	for i, v := range cnt {
		for v > 0 && v <= half && oddIdx < n {
			ans[oddIdx] = byte('a' + i)
			oddIdx += 2
			v--
		}
		for v > 0 {
			ans[evenIdx] = byte('a' + i)
			evenIdx += 2
			v--
		}
	}

	return string(ans)
}
