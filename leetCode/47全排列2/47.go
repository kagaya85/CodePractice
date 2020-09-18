package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}

	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) (ans [][]int) {
	// numMap := make(map[int]bool)
	sort.Ints(nums)
	var tmp []int
	flags := make([]bool, len(nums))

	var dfs func(pos int)
	dfs = func(pos int) {
		if pos >= len(nums) {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}

		for i, n := range nums {
			if flags[i] || i > 0 && n == nums[i-1] && !flags[i-1] {
				continue
			}

			tmp = append(tmp, n)
			flags[i] = true

			dfs(pos + 1)

			flags[i] = false
			tmp = tmp[:len(tmp)-1]

		}

	}

	dfs(0)
	return
}
