package main

import "sort"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	sumarr := make([]int, 0, m*n)
	presum := make([][]int, m+1)
	for i := range presum {
		presum[i] = make([]int, n+1)
	}
	for i, row := range matrix {
		for j, val := range row {
			presum[i+1][j+1] = presum[i+1][j] ^ presum[i][j+1] ^ presum[i][j] ^ val
			sumarr = append(sumarr, presum[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sumarr)))
	return sumarr[k-1]
}
