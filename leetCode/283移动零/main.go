package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	n := len(nums)
	p, q := 0, 0
	for q < n {
		if nums[q] == 0 {
			q++
		} else if nums[p] != 0 {
			p++
			if p > q {
				q = p
			}
		} else {
			nums[p], nums[q] = nums[q], nums[p]
			p++
			q++
		}
	}

}

func moveZeroes2(nums []int) {
	n := len(nums)
	p, q := 0, 0
	for q < n {
		if nums[q] != 0 {
			nums[p], nums[q] = nums[q], nums[p]
			p++
		}
		q++
	}
}
