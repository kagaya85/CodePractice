package main

func singleNumber(nums []int) int {
	a, b := 0, 0 // a, b分别表示高位和和低位
	for _, num := range nums {
		a, b = ^a&b&num|a&^b&^num, ^a&(b^num)
	}
	return b
}
