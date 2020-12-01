package main

import "sort"

func main() {

}

func searchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}
