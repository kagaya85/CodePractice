package main

import (
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	h, w := len(matrix), len(matrix[0])
	// 找target所在行的后一行
	row := sort.Search(h, func(i int) bool {
		return matrix[i][0] > target
	}) - 1

	if row < 0 {
		return false
	}

	col := sort.SearchInts(matrix[row], target)

	return col < w && matrix[row][col] == target
}
