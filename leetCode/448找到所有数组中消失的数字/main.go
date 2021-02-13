package main

func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)

	for _, v := range nums {
		idx := (v - 1) % n
		nums[idx] += n
	}

	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}
