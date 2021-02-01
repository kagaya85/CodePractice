package main

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	mapB := make(map[int]struct{}, len(B))
	for _, v := range A {
		sumA += v
	}
	for _, v := range B {
		sumB += v
		mapB[v] = struct{}{}
	}

	diff := (sumA - sumB) / 2
	for _, va := range A {
		vb := va - diff
		if _, ok := mapB[vb]; ok {
			return []int{va, vb}
		}
	}

	return []int{}
}
