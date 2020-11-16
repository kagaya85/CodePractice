package main

import "sort"

func main() {

}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		x, y := people[i], people[j]
		return x[0] < y[0] || x[0] == y[0] && x[1] > y[1]
	})

	ans := make([][]int, len(people))

	for _, p := range people {
		free := p[1] + 1
		for i := range ans {
			if ans[i] == nil {
				free--
				if free == 0 {
					ans[i] = p
					break
				}
			}
		}
	}
	return ans
}
