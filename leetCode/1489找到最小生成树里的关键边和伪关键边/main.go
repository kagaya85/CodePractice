package main

import (
	"sort"
)

type unionFind struct {
	parent []int
	size   []int
}

const (
	notInMST    = -1
	notCritical = 0
	critical    = 1
)

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	edgeType := make([]int, m)

	// 为边增加编号id
	for i := range edges {
		edges[i] = append(edges[i], i)
	}

	// 按权值排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	type neighbor struct {
		to, edgeID int
	}

	graph := make([][]neighbor, n)
	dfn := make([]int, n) // 遍历时间戳
	timestamp := 0
	var tarjan func(int, int) int
	tarjan = func(v, pid int) int {
		timestamp++
		dfn[v] = timestamp
		lowV := timestamp
		for _, e := range graph[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := tarjan(w, e.edgeID)
				if lowW > dfn[v] {
					edgeType[e.edgeID] = critical
				}
				lowV = min(lowV, lowW)
			} else if e.edgeID != pid {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}

	uf := newUnionFind(n)
	for i := 0; i < m; {
		vs := []int{}
		// 权值相同的边分为一组，建图，然后用 Tarjan 算法找桥边
		for weight := edges[i][2]; i < m && edges[i][2] == weight; i++ {
			e := edges[i]
			v, w, edgeID := uf.find(e[0]), uf.find(e[1]), e[3]
			if v != w {
				graph[v] = append(graph[v], neighbor{w, edgeID})
				graph[w] = append(graph[w], neighbor{v, edgeID})
				vs = append(vs, v, w)
			} else {
				edgeType[edgeID] = notInMST
			}
		}

		for _, v := range vs {
			if dfn[v] == 0 {
				tarjan(v, -1)
			}
		}

		for j := 0; j < len(vs); j += 2 {
			v, w := vs[j], vs[j+1]
			uf.union(v, w)
			graph[v] = nil
			graph[w] = nil
			dfn[v] = 0
			dfn[w] = 0
		}
	}

	var criticalEdges, notCriticalEdges []int
	for i, tp := range edgeType {
		if tp == notCritical {
			notCriticalEdges = append(notCriticalEdges, i)
		} else if tp == critical {
			criticalEdges = append(criticalEdges, i)
		}
	}

	return [][]int{criticalEdges, notCriticalEdges}
}

func min(arr ...int) (res int) {
	res = arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return
}
