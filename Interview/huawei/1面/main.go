package main

import (
	"fmt"
	"math"
)

// 非负整数数组，整数，m个非空连续子数组，设计算法使m个子数组各自和最大值最小

// [7,2,5,10,8], m = 2, 18

func main() {
	arr := []int{7, 2, 5, 10, 8}
	m := 2

	fmt.Println(SubArray(arr, m))
}

func SubArray(arr []int, m int) (ans int) {
	n := len(arr)
	minSum := math.MaxInt

	var dfs func(int, int, int)
	dfs = func(start, cnt, curMax int) {
		if cnt == m-1 {
			sum := 0
			for i := start; i < n; i++ {
				sum += arr[i]
			}
			minSum = min(minSum, max(curMax, sum))
			return
		}

		sum := 0
		for cur := start; cur < n-m+cnt; cur++ {
			sum += arr[cur]
			dfs(cur+1, cnt+1, max(curMax, sum))
		}
	}

	sum := 0
	for start := 0; start < n-m+1; start++ {
		sum += arr[start]
		dfs(start+1, 1, sum)
	}
	return minSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
