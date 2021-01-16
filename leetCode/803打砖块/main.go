package main

func hitBricks(grid [][]int, hits [][]int) []int {
	rows, cols := len(grid), len(grid[0])

	newGrid := make([][]int, rows)
	for i := range newGrid {
		newGrid[i] = make([]int, cols)
		for j := range newGrid[i] {
			newGrid[i][j] = grid[i][j]
		}
	}

	inArea := func(x, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	getIndex := func(x, y int) int {
		return x*cols + y
	}

	for _, hit := range hits {
		newGrid[hit[0]][hit[1]] = 0
	}

	roof := rows * cols
	uf := NewUnionFind(roof + 1)

	for i, v := range newGrid[0] {
		if v == 1 {
			uf.Union(i, roof)
		}
	}

	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if newGrid[i][j] == 1 {
				if newGrid[i-1][j] == 1 {
					uf.Union(getIndex(i-1, j), getIndex(i, j))
				}
				if j > 0 && newGrid[i][j-1] == 1 {
					uf.Union(getIndex(i, j-1), getIndex(i, j))
				}
			}
		}
	}

	hitLen := len(hits)
	res := make([]int, hitLen)

	for i := hitLen - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]

		if grid[x][y] == 0 {
			continue
		}

		origin := uf.GetSize(roof)

		if x == 0 { // 如果是补回的砖块，链接到屋顶
			uf.Union(y, roof)
		}

		for _, direct := range [...][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			newX, newY := x+direct[0], y+direct[1]

			if inArea(newX, newY) && newGrid[newX][newY] == 1 {
				uf.Union(getIndex(x, y), getIndex(newX, newY))
			}
		}

		current := uf.GetSize(roof)
		res[i] = max(0, current-origin-1)
		newGrid[x][y] = 1
	}

	return res
}

func max(arr ...int) (res int) {
	res = arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(from, to int) {
	rootf := uf.Find(from)
	roott := uf.Find(to)

	if rootf == roott {
		return
	}

	uf.parent[rootf] = roott
	uf.size[roott] += uf.size[rootf]
}

func (uf UnionFind) GetSize(x int) int {
	return uf.size[uf.Find(x)]
}
