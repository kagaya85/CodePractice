package main

func sumOddLengthSubarrays(arr []int) (ans int) {
	n := len(arr)
	presum := make([]int, n+1)
	for i := range arr {
		presum[i+1] = presum[i] + arr[i]
	}
	for start := 0; start < n; start++ {
		for length := 1; start+length <= n; length += 2 {
			end := start + length
			ans += presum[end] - presum[start]
		}
	}
	return
}
