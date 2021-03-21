package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	iscol0 := false // 第0列是否需要清0

	for _, r := range matrix {
		if r[0] == 0 {
			iscol0 = true
		}
		for i := 1; i < n; i++ {
			if r[i] == 0 {
				r[0] = 0         // 使用第0列记录需要清0的行
				matrix[0][i] = 0 // 使用第0行记录需要清0的列
			}
		}
	}

	for i := m - 1; i >= 0; i-- { // 从下往上遍历，最后处理第0行
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if iscol0 { // 最后处理第0列
			matrix[i][0] = 0
		}
	}
}
