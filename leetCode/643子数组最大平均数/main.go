package main

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0

	for _, v := range nums[:k] {
		sum += v
	}

	maxAvg := float64(sum) / float64(k)
	for i := k; i < n; i++ {
		sum = sum + nums[i] - nums[i-k]
		maxAvg = max(maxAvg, float64(sum)/float64(k))
	}
	return maxAvg
}

func max(arr ...float64) float64 {
	res := arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
