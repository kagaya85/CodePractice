package main

func findErrorNums(nums []int) (ans []int) {
	cnt := map[int]int{}
	for _, n := range nums {
		cnt[n]++
	}
	ans = make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if v := cnt[i]; v == 2 {
			ans[0] = i
		} else if v == 0 {
			ans[1] = i
		}
	}
	return ans
}
