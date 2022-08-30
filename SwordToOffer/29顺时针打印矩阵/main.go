package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(SpiralMatrix(matrix))
}

func SpiralMatrix(matrix [][]int) []int {
	// write code here
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	result := make([]int, m*n)

	left, top, right, bottom := 0, 0, n, m
	idx := 0
	for left < right && top < bottom {
		for col, row := left, top; col < right; col++ {
			result[idx] = matrix[row][col]
			idx++
		}
		for col, row := right-1, top+1; row < bottom; row++ {
			result[idx] = matrix[row][col]
			idx++
		}
		if top+1 < bottom && left+1 < right {
			for col, row := right-2, bottom-1; col > left; col-- {
				result[idx] = matrix[row][col]
				idx++
			}
			for col, row := left, bottom-1; row > top; row-- {
				result[idx] = matrix[row][col]
				idx++
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return result
}
