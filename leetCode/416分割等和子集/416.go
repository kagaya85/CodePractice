package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}

	sum, max := 0, 0
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if target < max {
		return false
	}

	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	dp[0][nums[0]] = true

	for i := 1; i < length; i++ {
		dp[i][0] = true
		n := nums[i]
		for j := 1; j <= target; j++ {
			if n > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-n]
			}
		}
	}

	return dp[length-1][target]
}
