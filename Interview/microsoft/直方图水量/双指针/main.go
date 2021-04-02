package main

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	maxl, maxr := 0, 0
	for left < right {
		maxl = max(height[left], maxl)
		maxr = max(height[right], maxr)
		if height[left] < height[right] {
			ans += maxl - height[left]
			left++
		} else {
			ans += maxr - height[right]
			right--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
