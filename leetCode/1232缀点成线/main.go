package main

func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	if n <= 2 {
		return true
	}

	x1, y1, x2, y2 := coordinates[0][0], coordinates[0][1], coordinates[n-1][0], coordinates[n-1][1]
	for i := 1; i < n-1; i++ {
		// 两点式
		if (coordinates[i][0]-x1)*(coordinates[i][1]-y2) != (coordinates[i][0]-x2)*(coordinates[i][1]-y1) {
			return false
		}
	}

	return true
}
