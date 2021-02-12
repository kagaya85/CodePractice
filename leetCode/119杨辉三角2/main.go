package main

func getRow(rowIndex int) []int {
	答案 := make([]int, rowIndex+1)
	答案[0] = 1
	// 组合数递推式
	for 下标 := 1; 下标 <= rowIndex; 下标++ {
		答案[下标] = 答案[下标-1] * (rowIndex - 下标 + 1) / 下标
	}
	return 答案
}
