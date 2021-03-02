package main

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}

	n := len(matrix[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	for i, row := range matrix {
		sum[i+1] = make([]int, n+1)
		for j, v := range row {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
}
