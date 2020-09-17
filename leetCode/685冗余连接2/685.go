package main

func main() {

}

func findRedundantDirectedConnection(edges [][]int) []int {
	nodesNum := len(edges)

	uf := newUnionFind(nodesNum + 1)
	parent := make([]int, nodesNum+1)

	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to {
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) {
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}

	if conflictEdge == nil {
		return cycleEdge
	}

	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}

	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)

	for i := 0; i < n; i++ {
		ancestor[i] = i
	}

	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		return uf.find(uf.ancestor[x])
	} else {
		return uf.ancestor[x]
	}
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
