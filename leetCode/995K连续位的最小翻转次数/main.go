package main

func minKBitFlips(A []int, K int) (ans int) {
	fliped := 0
	n := len(A)
	for i, v := range A {

		if i >= K && A[i-K] > 1 { // 判断是否需要修正fliped
			fliped ^= 1
		}

		if v == fliped { // 判断是否需要翻转
			if i+K > n {
				return -1
			}
			A[i] += 2
			fliped ^= 1
			ans++
		}
	}
	return
}
