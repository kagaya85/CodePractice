package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	fmt.Println(subsets(nums))
}

func subsets(nums []int) (ans [][]int) {

	sub := []int{}

	var dfs func(start, n int)
	dfs = func(start, n int) {
		// fmt.Printf("start: %d, n: %d\n", start, n)
		if len(nums)-start < n {
			return
		}

		if n == 0 {
			ans = append(ans, append([]int(nil), sub...))
			return
		}

		for i := start; i < len(nums); i++ {
			sub = append(sub, nums[i])
			// fmt.Printf("before sub: %v, start: %d, i: %d\n", sub, start, i)
			dfs(i+1, n-1)
			sub = sub[:len(sub)-1]
		}
	}

	for n := 0; n <= len(nums); n++ {
		dfs(0, n)
	}

	return
}
