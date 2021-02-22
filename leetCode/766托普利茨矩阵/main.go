package main

func isToeplitzMatrix(matrix [][]int) bool {
	w, h := len(matrix[0]), len(matrix)
	for idx, v := range matrix[0] {
		i, j := 1, idx+1
		for i < h && j < w {
			if matrix[i][j] != v {
				return false
			}
			i++
			j++
		}
	}

	for idx, line := range matrix {
		v := line[0]
		i, j := idx+1, 1
		for i < h && j < w {
			if matrix[i][j] != v {
				return false
			}
			i++
			j++
		}
	}
	return true
}
