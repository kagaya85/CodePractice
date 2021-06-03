package main

func findMaxLength(nums []int) (ans int) {
	preSumIdx := map[int]int{0: -1}
	sum := 0
	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}
		if preIdx, has := preSumIdx[sum]; has {
			ans = max(ans, i-preIdx)
		} else {
			preSumIdx[sum] = i
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
