package main

import "sort"

func guessNumber(n int) int {
	return sort.Search(n, func(i int) bool {
		return guess(i) <= 0
	})
}

func guess(num int) int {
	return 0
}
