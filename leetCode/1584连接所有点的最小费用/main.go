package main

import "sort"

type edge struct {
	x, y, dist int
}

func minCostConnectPoints(points [][]int) (ans int) {
	n := len(points)
	edges := []edge{}

	for i, p := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, distance(p, points[j])})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
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

	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, e := range edges {
		if find(e.x) != find(e.y) {
			union(e.x, e.y)
			ans += e.dist
		}
	}

	return
}

func distance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
