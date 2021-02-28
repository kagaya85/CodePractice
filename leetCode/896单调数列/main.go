package main

const (
	none int = iota
	up
	down
)

func isMonotonic(A []int) bool {
	if len(A) < 3 {
		return true
	}

	var status int
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			if status == down {
				return false
			}
			status = up
		} else if A[i-1] > A[i] {
			if status == up {
				return false
			}
			status = down
		}
	}

	return true
}
