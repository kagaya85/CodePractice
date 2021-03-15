package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	left, right, top, bottom := 0, n-1, 0, m-1
	idx := 0
	for left <= right && top <= bottom {
		// 顺时针遍历
		for i := left; i < right; i++ {
			ans[idx] = matrix[top][i]
			idx++
		}
		for i := top; i <= bottom; i++ {
			ans[idx] = matrix[i][right]
			idx++
		}
		// 防止重复遍历
		if left == right || top == bottom {
			break
		}

		for i := right - 1; i > left; i-- {
			ans[idx] = matrix[bottom][i]
			idx++
		}
		for i := bottom; i > top; i-- {
			ans[idx] = matrix[i][left]
			idx++
		}

		left++
		right--
		top++
		bottom--
	}
	return ans
}
