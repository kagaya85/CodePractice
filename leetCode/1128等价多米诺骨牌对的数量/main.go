package main

func numEquivDominoPairs(dominoes [][]int) (ans int) {
	numCnt := make([]int, 100)
	for _, v := range dominoes {
		var idx int
		if v[0] < v[1] {
			idx = v[0]*10 + v[1]
		} else {
			idx = v[1]*10 + v[0]
		}
		ans += numCnt[idx]
		numCnt[idx]++
	}
	return
}
