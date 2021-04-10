package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, div := range []int{2, 3, 5} {
		for n%div == 0 {
			n /= div
		}
	}
	return n == 1
}
