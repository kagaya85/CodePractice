package main

func longestSubarray(nums []int, limit int) int {
	n := len(nums)
	var left, right int
	var maxq, minq []int

	for right = 0; right < n; right++ {
		v := nums[right]
		// right value append to priority queue
		for len(maxq) > 0 && maxq[len(maxq)-1] < v {
			maxq = maxq[:len(maxq)-1]
		}
		maxq = append(maxq, v)
		for len(minq) > 0 && minq[len(minq)-1] > v {
			minq = minq[:len(minq)-1]
		}
		minq = append(minq, v)
		// check min max delta

		if len(minq) > 0 && len(maxq) > 0 && maxq[0]-minq[0] > limit {
			if nums[left] == minq[0] {
				minq = minq[1:]
			}
			if nums[left] == maxq[0] {
				maxq = maxq[1:]
			}
			left++
		}
	}

	return right - left
}
