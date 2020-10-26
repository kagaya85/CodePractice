package main

import (
	"fmt"
	"sort"
)

type pair struct {
	val int
	idx int
}

func main() {

	nums := []int{8, 1, 2, 2, 3}

	fmt.Println(smallerNumbersThanCurrent(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
	data := make([]pair, len(nums))
	ans := make([]int, len(nums))

	for i, v := range nums {
		data[i] = pair{v, i}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].val < data[j].val
	})

	pre := 0
	for i, p := range data {
		if i > 0 && data[i-1].val != p.val {
			pre = i
		}
		ans[p.idx] = pre
	}

	return ans
}
