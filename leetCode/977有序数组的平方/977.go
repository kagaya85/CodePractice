package main

func main() {

}

func sortedSquares(A []int) (ans []int) {
	n := len(A)
	lastNegIndex := 0

	for i := 0; i < n && A[i] < 0; i++ {
		lastNegIndex = i
	}

	left, right := lastNegIndex, lastNegIndex+1

	for left >= 0 && right < n {
		if A[left]*A[left] <= A[right]*A[right] {
			ans = append(ans, A[left]*A[left])
			left--
		} else {
			ans = append(ans, A[right]*A[right])
			right++
		}
	}

	for left >= 0 {
		ans = append(ans, A[left]*A[left])
		left--
	}

	for right < n {
		ans = append(ans, A[right]*A[right])
		right++
	}

	return
}
