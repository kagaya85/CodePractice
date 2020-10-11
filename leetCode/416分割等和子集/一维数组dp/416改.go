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

	// 压缩dp数组为一维数组
	dp := make([]bool, target+1)
	dp[0] = true
	dp[nums[0]] = true

	for i := 1; i < length; i++ {
		n := nums[i]
		for j := target; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
		}
	}

	return dp[target]
}
