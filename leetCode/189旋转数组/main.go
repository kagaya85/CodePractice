package main

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse := func(arr []int, from, to int) {
		to--
		for from < to {
			arr[from], arr[to] = arr[to], arr[from]
			from++
			to--
		}
	}

	reverse(nums, 0, n)
	reverse(nums, 0, k)
	reverse(nums, k, n)
}
