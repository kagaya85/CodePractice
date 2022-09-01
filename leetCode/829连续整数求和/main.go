package main

func consecutiveNumbersSum(n int) int {
	count, sum := 0, 1
	for k := 1; sum <= n; k++ {
		if (n-sum)%k == 0 {
			count++
		}
		sum += k + 1
	}
	return count
}
