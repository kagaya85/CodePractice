package main

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	res := make([][]int, m)

	for i := range res {
		res[i] = make([]int, n)
	}

	for i := range matrix {
		for j, v := range matrix[i] {
			res[j][i] = v
		}
	}
	return res
}
