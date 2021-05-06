package main

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i, e := range encoded {
		res[i+1] = e ^ res[i]
	}
	return res
}
