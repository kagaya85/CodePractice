package main

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func quickPow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for n > 0 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
		n >>= 1
	}
	return ret
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	res := quickPow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}
