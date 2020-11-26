package main

import "fmt"

func main() {
	nums := []int{3, 6, 9, 1}

	fmt.Println(maximumGap(nums))
}

type bucket struct {
	min int
	max int
}

func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	maxVal := max(nums...)
	minVal := min(nums...)
	bktDist := max(1, (maxVal-minVal)/(n-1))
	bktNum := (maxVal-minVal)/bktDist + 1
	buckets := make([]bucket, bktNum)

	for i := range buckets {
		buckets[i] = bucket{-1, -1}
	}

	for _, v := range nums {
		idx := (v - minVal) / bktDist
		if buckets[idx].min == -1 {
			buckets[idx].min = v
			buckets[idx].max = v
		} else {
			buckets[idx].min = min(buckets[idx].min, v)
			buckets[idx].max = max(buckets[idx].max, v)
		}
	}

	prev := -1
	for i, b := range buckets {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			ans = max(ans, b.min-buckets[prev].max)
		}
		prev = i
	}

	return
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

func max(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
