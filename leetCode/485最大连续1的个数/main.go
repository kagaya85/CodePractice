package main

func findMaxConsecutiveOnes(nums []int) (ans int) {

	count := 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			ans = max(ans, count)
			count = 0
		}
	}
	ans = max(ans, count)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
