package main

type unionFind struct {
	parent   []int
	setCount int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	setCount := n
	for i := range parent {
		parent[i] = i
	}
	return &unionFind{
		parent:   parent,
		setCount: setCount,
	}
}

func (uf *unionFind) union(x, y int) bool {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return false
	}
	uf.parent[px] = py
	uf.setCount--
	return true
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	delNum := len(edges)
	alice, bob := newUnionFind(n), newUnionFind(n)

	for _, e := range edges {
		if e[0] == 3 {
			x, y := e[1]-1, e[2]-1
			if alice.find(x) != alice.find(y) {
				alice.union(x, y)
				bob.union(x, y)
				delNum--
			}
		}
	}

	uf := [2]*unionFind{alice, bob}
	for _, e := range edges {
		x, y := e[1]-1, e[2]-1
		if idx := e[0]; idx < 3 && uf[idx-1].union(x, y) {
			delNum--
		}
	}

	if alice.setCount > 1 || bob.setCount > 1 {
		return -1
	}

	return delNum
}
