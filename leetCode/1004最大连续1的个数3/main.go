package main

func longestOnes(A []int, K int) int {
	zeroCnt, n := 0, len(A)
	left, right := 0, 0

	for ; right < n; right++ {
		if A[right] == 0 {
			zeroCnt++
		}
		if zeroCnt > K {
			if A[left] == 0 {
				zeroCnt--
			}
			left++
		}
	}

	return right - left
}
