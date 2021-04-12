package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		return combine(x, y) > combine(y, x)
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, n := range nums {
		ans = append(ans, strconv.Itoa(n)...)
	}
	return string(ans)
}

func combine(x, y int) int {
	t := 10
	for t <= y {
		t *= 10
	}
	return x*t + y
}
