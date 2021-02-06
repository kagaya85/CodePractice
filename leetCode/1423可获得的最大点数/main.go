package main

func maxScore(cardPoints []int, k int) int {
	sum := 0
	n, l := len(cardPoints), len(cardPoints)-k
	for _, v := range cardPoints[:l] {
		sum += v
	}

	minSum := sum
	for i := l; i < n; i++ {
		sum -= cardPoints[i-l]
		sum += cardPoints[i]
		minSum = min(minSum, sum)
	}

	tot := 0
	for _, v := range cardPoints {
		tot += v
	}

	return tot - minSum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
