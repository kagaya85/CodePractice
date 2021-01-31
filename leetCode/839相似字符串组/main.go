package main

func numSimilarGroups(strs []string) int {
	n := len(strs)
	parent := make([]int, n)
	unionCnt := n

	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		px, py := find(x), find(y)
		if px == py {
			return
		}
		parent[px] = py
		unionCnt--
	}

	isUnion := func(x, y int) bool {
		if find(x) == find(y) {
			return true
		}
		return false
	}

	for i := range strs {
		for j := i + 1; j < n; j++ {
			if !isUnion(i, j) && isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}

	return unionCnt
}

func isSimilar(x, y string) bool {
	cnt := 0
	for i := range x {
		if x[i] != y[i] {
			cnt++
		}
		if cnt > 2 {
			return false
		}
	}

	return true
}
