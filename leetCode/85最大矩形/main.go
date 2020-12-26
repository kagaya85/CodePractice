package main

func maximalRectangle(matrix [][]byte) (ans int) {
	yl := len(matrix)
	if yl == 0 {
		return
	}
	xl := len(matrix[0])
	// 统计每行的连续1个数
	cnt := make([][]int, yl)
	for i, row := range matrix {
		cnt[i] = make([]int, xl)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				cnt[i][j] = 1
			} else {
				cnt[i][j] = cnt[i][j-1] + 1
			}
		}
	}

	// 对每一列搜索最大矩形
	for x := 0; x < xl; x++ {
		up := make([]int, yl)
		down := make([]int, yl)
		stk := []int{}

		// 计算向上扩展的边界
		for y := 0; y < yl; y++ {
			// 出栈
			for len(stk) > 0 && cnt[stk[len(stk)-1]][x] >= cnt[y][x] {
				stk = stk[:len(stk)-1]
			}
			up[y] = -1
			if len(stk) > 0 {
				up[y] = stk[len(stk)-1]
			}
			stk = append(stk, y)
		}
		stk = stk[:0]
		// 计算向下扩展的边界
		for y := yl - 1; y >= 0; y-- {
			// 出栈
			for len(stk) > 0 && cnt[stk[len(stk)-1]][x] >= cnt[y][x] {
				stk = stk[:len(stk)-1]
			}
			down[y] = yl
			if len(stk) > 0 {
				down[y] = stk[len(stk)-1]
			}
			stk = append(stk, y)
		}
		// 计算该列每一层所能上下扩展的最大矩形
		for y, row := range cnt {
			h := down[y] - up[y] - 1
			area := h * row[x]
			ans = max(ans, area)
		}
	}
	return
}

func max(arr ...int) (res int) {
	res = arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return
}
