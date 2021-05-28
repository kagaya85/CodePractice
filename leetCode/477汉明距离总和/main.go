package main

func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i <= 30; i++ {
		c := 0
		for _, num := range nums {
			c += num >> i & 1
		}
		ans += c * (n - c)
	}
	return
}
