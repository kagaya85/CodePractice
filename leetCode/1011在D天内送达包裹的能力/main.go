package main

import "sort"

func shipWithinDays(weights []int, D int) int {
	l, r := 0, 0
	for _, w := range weights {
		if w > l {
			l = w
		}
		r += w
	}
	return l + sort.Search(r-l, func(x int) bool {
		x += l
		day, sum := 1, 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}
