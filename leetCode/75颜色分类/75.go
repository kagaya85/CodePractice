package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	p0, p1 := 0, 0

	for i, n := range nums {
		if n == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		} else if n == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		}
	}
}
