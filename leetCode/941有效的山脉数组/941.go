package main

import "fmt"

func main() {

	A := []int{3, 2, 1}

	fmt.Println(validMountainArray(A))
}

func validMountainArray(A []int) bool {
	l := len(A)
	if l < 3 || A[0] > A[1] {
		return false
	}

	flag := 0
	pre := A[0]
	for i := 1; i < l; i++ {
		cur := A[i]
		if flag == 0 && cur < pre {
			flag = 1
		} else if flag == 1 && cur > pre || cur == pre {
			return false
		}
		pre = cur
	}

	if flag == 1 {
		return true
	}
	return false
}
