package main

func main() {

}

func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]%2 != 0 {
			for A[j]%2 != 0 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
