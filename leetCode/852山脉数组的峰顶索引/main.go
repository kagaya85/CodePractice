package main

import "sort"

func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] > arr[i+1]
	})
}
