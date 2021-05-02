package main

func leastBricks(wall [][]int) int {
	edgeCount := make(map[int]int)
	for _, row := range wall {
		distL := 0
		for _, length := range row[:len(row)-1] {
			distL += length
			edgeCount[distL]++
		}
	}

	max := 0
	for _, v := range edgeCount {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}
