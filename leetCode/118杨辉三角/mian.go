package main

import "fmt"

func main() {
	fmt.Println(generate(0))
}

func generate(rows int) (ans [][]int) {
	nextLine := func(prevLine []int) []int {
		n := len(prevLine)
		res := make([]int, n+1)
		res[0], res[n] = 1, 1
		for i := 1; i < n; i++ {
			res[i] = prevLine[i-1] + prevLine[i]
		}
		return res
	}

	if rows >= 1 {
		prev := []int{1}
		ans = append(ans, prev)
		for i := 1; i < rows; i++ {
			next := nextLine(prev)
			ans = append(ans, next)
			prev = next
		}
	}

	return ans
}
