package main

func largeGroupPositions(s string) (ans [][]int) {

	n := len(s)
	if n < 3 {
		return ans
	}

	start := 0
	count := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			if count >= 3 {
				ans = append(ans, []int{start, i - 1})
			}
			count = 1
			start = i
		}
	}

	if count >= 3 {
		ans = append(ans, []int{start, n - 1})
	}

	return ans
}
