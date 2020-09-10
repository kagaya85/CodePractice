package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{7, 3, 2}
	target := 18

	var ans [][]int = combinationSum(candidates, target)

	fmt.Println(ans)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var ans [][]int

	backtrack(0, 0, make([]int, 0), candidates, target, &ans)

	return ans
}

func backtrack(i int, tmpSum int, tmp []int, candidates []int, target int, ans *[][]int) {

	l := len(candidates)
	fmt.Printf("i: %d, tmpSum: %d, tmp: %v\n", i, tmpSum, tmp)

	if tmpSum > target || i == l {
		return
	}
	if tmpSum == target {
		subAns := make([]int, len(tmp))
		copy(subAns, tmp)
		*ans = append(*ans, subAns)
		fmt.Printf("add subAns: %v, ans: %v\n", subAns, *ans)
		return
	}

	for j := i; j < l; j++ {
		if tmpSum+candidates[j] > target {
			break
		}
		backtrack(j, tmpSum+candidates[j], append(tmp, candidates[j]), candidates, target, ans)
	}
}
