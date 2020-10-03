package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, n := range nums {
		if index, ok := hashMap[target-n]; ok {
			return []int{index, i}
		}
		hashMap[n] = i
	}

	return nil
}
