package main

func main() {

}

func fourSumCount(A, B, C, D []int) (ans int) {
	sum := make(map[int]int)

	for _, a := range A {
		for _, b := range B {
			sum[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			ans += sum[-c-d]
		}
	}

	return
}
