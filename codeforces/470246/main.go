package main

import "fmt"

/*
Theofanis has a string s1s2…sn and a character c. He wants to make all characters of the string equal to c using the minimum number of operations.

In one operation he can choose a number x (1≤x≤n) and for every position i,
where i is not divisible by x, replace si with c.

Find the minimum number of operations required to make all the characters equal to c and the x-s that he should use in his operations.
*/
func main() {
	var t int
	fmt.Scan(&t)

	var n int
	var c byte
	var s string
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &c, &s)
		m, xs := minOperations(n, c, s)
		fmt.Println(m)
		if m > 0 {
			for _, x := range xs {
				fmt.Printf("%d ", x)
			}
			fmt.Println()
		}
	}
}

func minOperations(n int, c byte, s string) (int, []int) {
	var count int
	for i := 0; i < n; i++ {
		if s[i] != c {
			count++
		}
	}

	if count == 0 {
		return 0, nil
	}

	if count == n {
		return 1, []int{n}
	}

	if n%count == 0 {
		return 1, []int{n / count}
	}

	return 2, []int{n - count, n}
}
