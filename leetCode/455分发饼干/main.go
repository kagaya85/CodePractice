package main

import "sort"

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)

	gl, sl := len(g), len(s)
	for i, j := 0, 0; i < gl && j < sl; i++ {
		for j < sl && g[i] > s[j] {
			j++
		}
		if j < sl {
			ans++
			j++
		}
	}

	return
}
