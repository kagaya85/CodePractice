package main

func main() {

}

func insert(intervals [][]int, newInterval []int) (ans [][]int) {

	left, right := newInterval[0], newInterval[1]
	needMerge := true
	for _, itv := range intervals {
		if itv[0] > right {
			if needMerge {
				ans = append(ans, []int{left, right})
				needMerge = false
			}
			ans = append(ans, itv)
		} else if itv[1] < left {
			ans = append(ans, itv)
		} else {
			left = min(left, itv[0])
			right = max(right, itv[1])
		}

	}

	if needMerge {
		ans = append(ans, []int{left, right})
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
