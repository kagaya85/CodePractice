package main

import "sort"

type edge struct {
	x, y, cost int
}

func minimumEffortPath(heights [][]int) int {
	rows := len(heights)
	cols := len(heights[0])
	n := rows * cols

	edges := make([]edge, 0, 2*n-rows-cols)
	for i := range heights {
		for j := range heights[0] {
			idx := i*cols + j
			if i < rows-1 {
				down := idx + cols
				edges = append(edges, edge{idx, down, abs(heights[i][j] - heights[i+1][j])})
			}
			if j < cols-1 {
				right := idx + 1
				edges = append(edges, edge{idx, right, abs(heights[i][j] - heights[i][j+1])})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

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

	union := func(x, y int) {
		parent[find(x)] = find(y)
	}

	for _, e := range edges {
		union(e.x, e.y)
		if find(0) == find(n-1) {
			return e.cost
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
