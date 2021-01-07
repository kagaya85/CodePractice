package main

func findCircleNum(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parents := make([]int, n)

	for i := range parents {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	union := func(from, to int) {
		parents[find(from)] = parents[find(to)]
	}

	for i := range isConnected {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	for i := range parents {
		if i == parents[i] {
			ans++
		}
	}

	return
}
