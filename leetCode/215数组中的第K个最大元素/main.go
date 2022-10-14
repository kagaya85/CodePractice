package main

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l int, r int, index int) int {
	i := partition(nums, l, r)
	if i == index {
		return nums[i]
	}

	if i < index {
		return quickSelect(nums, i+1, r, index)
	} else {
		return quickSelect(nums, l, i-1, index)
	}
}

func partition(nums []int, l int, r int) int {
	ri := rand.Int()%(r-l+1) + l
	nums[r], nums[ri] = nums[ri], nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < nums[r] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
