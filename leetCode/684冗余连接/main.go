package main

func findRedundantConnection(edges [][]int) []int {

	n := len(edges)
	ancestor := make([]int, n+1)
	for i := range ancestor {
		ancestor[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if ancestor[x] != x {
			ancestor[x] = find(ancestor[x])
		}
		return ancestor[x]
	}

	union := func(from, to int) {
		ancestor[find(from)] = find(to)
	}

	var cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if find(from) == find(to) {
			cycleEdge = edge
		} else {
			union(from, to)
		}
	}

	return cycleEdge
}
