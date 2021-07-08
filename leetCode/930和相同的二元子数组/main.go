package main

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	cnt := map[int]int{}
	sum := 0
	for _, num := range nums {
		cnt[sum]++
		sum += num
		ans += cnt[sum-goal]
	}
	return
}
