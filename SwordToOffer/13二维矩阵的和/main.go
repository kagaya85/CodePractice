type NumMatrix struct {
	rows int
	cols int
	sum  [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	if rows == 0 {
		return NumMatrix{}
	}
	cols := len(matrix[0])
	sum := make([][]int, rows)
	for i := range sum {
		sum[i] = make([]int, cols)
	}

	for i := range sum {
		for j := range sum[i] {
			if i == 0 && j > 0 {
				sum[i][j] = sum[i][j-1]
			} else if i > 0 && j == 0 {
				sum[i][j] = sum[i-1][j]
			} else if i > 0 && j > 0 {
				sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
			}
			sum[i][j] += matrix[i][j]
		}
	}

	return NumMatrix{
		rows: rows,
		cols: cols,
		sum:  sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := this.sum[row2][col2]
	if row1 > 0 {
		res -= this.sum[row1-1][col2]
	}
	if col1 > 0 {
		res -= this.sum[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		res += this.sum[row1-1][col1-1]
	}
	return res
}