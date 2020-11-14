package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}

	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := make(map[int]int)
	for i, v := range arr2 {
		rank[v] = i - len(arr2)
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		if r, has := rank[x]; has {
			x = r
		}
		if r, has := rank[y]; has {
			y = r
		}
		return x < y
	})
	return arr1
}
