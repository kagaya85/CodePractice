package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	numS := make([]int, 0, 20)

	for i := range tokens {
		s := tokens[i]

		switch s {
		case "+":
			l := len(numS)
			numS[l-2] += numS[l-1]
			numS = numS[:l-1]
		case "-":
			l := len(numS)
			numS[l-2] -= numS[l-1]
			numS = numS[:l-1]
		case "*":
			l := len(numS)
			numS[l-2] *= numS[l-1]
			numS = numS[:l-1]
		case "/":
			l := len(numS)
			numS[l-2] /= numS[l-1]
			numS = numS[:l-1]
		default:
			num, _ := strconv.Atoi(s)
			numS = append(numS, num)
		}
	}
	if len(numS) != 1 {
		fmt.Println("numS has many elems")
	}
	return numS[0]
}
