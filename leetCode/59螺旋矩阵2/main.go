package main

func generateMatrix(n int) (ans [][]int) {
	ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	idx := 1
	for left <= right && top <= bottom {
		for i := left; i < right; i++ {
			ans[top][i] = idx
			idx++
		}
		for i := top; i <= bottom; i++ {
			ans[i][right] = idx
			idx++
		}
		if left == right || top == bottom {
			break
		}
		for i := right - 1; i > left; i-- {
			ans[bottom][i] = idx
			idx++
		}
		for i := bottom; i > top; i-- {
			ans[i][left] = idx
			idx++
		}
		left++
		right--
		top++
		bottom--
	}

	return
}
