package main

func regionsBySlashes(grid []string) int {
	n := len(grid)
	count := 4 * n * n
	parent := make([]int, count)
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

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = parent[pb]
			count--
		}
	}

	//	---------
	//		1
	//	0		2
	//		3
	//	---------
	for i := range grid {
		for j := range grid {
			idx := (i*n + j) * 4
			if i < n-1 {
				// 和下面的区域连接
				down := ((i+1)*n + j) * 4
				union(idx+3, down+1)
			}

			if j < n-1 {
				// 和右边的区域连接
				right := (i*n + j + 1) * 4
				union(idx+2, right)
			}

			if grid[i][j] == '/' {
				union(idx, idx+1)
				union(idx+2, idx+3)
			} else if grid[i][j] == '\\' {
				union(idx+1, idx+2)
				union(idx, idx+3)
			} else {
				union(idx, idx+1)
				union(idx+1, idx+2)
				union(idx+2, idx+3)
			}
		}
	}

	return count
}
