package main

func deleteAndEarn(nums []int) int {
	maxN := 0
	for _, n := range nums {
		maxN = max(maxN, n)
	}
	sum := make([]int, maxN+1)
	for _, n := range nums {
		sum[n] += n
	}
	first, second := sum[0], max(sum[0], sum[1])
	for i := 2; i < len(sum); i++ {
		first, second = second, max(first+sum[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
