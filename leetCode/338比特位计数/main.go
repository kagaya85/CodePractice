package main

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 最低位
		ans[i] = ans[i>>1] + i&1
		// 最低为1的位
		// ans[i] = ans[i&(i-1)] + 1
	}
	return ans
}
