package main

import "fmt"

func main() {

	nums := []int{-2, 5, -1}
	lower, upper := -2, 2

	fmt.Println(countRangeSum(nums, lower, upper))
}

func countRangeSum(nums []int, lower int, upper int) int {

	var mergeSortCount func(arr []int) int
	mergeSortCount = func(arr []int) int {
		l := len(arr)
		if l <= 1 {
			return 0
		}

		a1 := append([]int(nil), arr[:l/2]...)
		a2 := append([]int(nil), arr[l/2:]...)

		count := mergeSortCount(a1) + mergeSortCount(a2)

		l, r := 0, 0
		for _, n := range a1 {
			for l < len(a2) && a2[l]-n < lower {
				l++
			}
			r = l
			for r < len(a2) && a2[r]-n <= upper {
				r++
			}
			count += r - l
		}

		// 归并排序
		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(a1) && (p2 >= len(a2) || a1[p1] <= a2[p2]) {
				arr[i] = a1[p1]
				p1++
			} else {
				arr[i] = a2[p2]
				p2++
			}
		}

		return count
	}

	prefixSum := make([]int, len(nums)+1)
	for i, n := range nums {
		prefixSum[i+1] = prefixSum[i] + n
	}

	return mergeSortCount(prefixSum)
}
