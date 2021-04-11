package main

func nthUglyNumber(n int) int {
	nums := make([]int, n+1)
	nums[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		n2, n3, n5 := nums[p2]*2, nums[p3]*3, nums[p5]*5
		nums[i] = min(n2, n3, n5)
		if nums[i] == n2 {
			p2++
		}
		if nums[i] == n3 {
			p3++
		}
		if nums[i] == n5 {
			p5++
		}
	}
	return nums[n]
}

func min(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}
	return res
}
