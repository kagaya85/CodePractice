package main

func subarraySum(arr []int, target int) (ans int) {
	cnt := map[int]int{0: 1}
	presum := 0
	for _, num := range arr {
		presum += num
		if _, has := cnt[presum-target]; has {
			ans += cnt[presum-target]
		}
		cnt[presum]++
	}
	return
}

func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
	n := len(matrix[0])
	for up := range matrix {
		presum := make([]int, n)
		for _, row := range matrix[up:] {
			for i, num := range row {
				presum[i] += num
			}
			ans += subarraySum(presum, target)
		}
	}
	return
}
