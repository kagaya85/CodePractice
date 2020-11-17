package main

import "fmt"

func main() {
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	maxDist := max(r0, R-1-r0) + max(c0, C-1-c0)
	buckets := make([][][]int, maxDist+1)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := abs(i-r0) + abs(j-c0)
			buckets[dist] = append(buckets[dist], []int{i, j})
		}
	}
	ans := make([][]int, 0, R*C)
	for _, b := range buckets {
		ans = append(ans, b...)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
