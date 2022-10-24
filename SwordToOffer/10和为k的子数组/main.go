func subarraySum(nums []int, k int) (ans int) {
	sum := map[int]int{0: 1}
	pre := 0

	for _, v := range nums {
		pre += v
		ans += sum[pre-k]
		sum[pre]++
	}

	return
}