package main

func makeConnected(n int, connections [][]int) int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	var extraNum int
	for _, conn := range connections {
		pa, pb := find(conn[0]), find(conn[1])
		if pa != pb {
			union(pa, pb)
		} else {
			extraNum++
		}
	}

	var connectNum int
	for i := range parent {
		if i == parent[i] {
			connectNum++
		}
	}

	if extraNum >= connectNum-1 {
		return connectNum - 1
	}
	return -1
}
