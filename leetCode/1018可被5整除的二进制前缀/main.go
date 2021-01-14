package main

func prefixesDivBy5(A []int) []bool {
	ans := make([]bool, len(A))
	remain := 0
	for i, v := range A {
		remain = (remain<<1 | v) % 5
		ans[i] = remain == 0
	}
	return ans
}
