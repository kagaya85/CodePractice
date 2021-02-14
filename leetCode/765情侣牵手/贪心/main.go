package main

func minSwapsCouples(row []int) (ans int) {
	n := len(row)
	// 反向索引
	pos := make([]int, n)
	for i, v := range row {
		pos[v] = i
	}

	for i := 0; i < n; i += 2 {
		if row[i] == row[i+1]^1 {
			continue
		}
		halfPos := pos[row[i]^1]
		row[i+1], row[halfPos] = row[halfPos], row[i+1]
		// 更新索引
		// pos[row[i+1]] = i + 1
		pos[row[halfPos]] = halfPos
		ans++
	}
	return
}
