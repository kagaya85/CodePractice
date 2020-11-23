package main

import "sort"

func main() {

}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrowIndex := points[0][1]
	ans := 1

	for _, p := range points {
		if p[0] > arrowIndex {
			arrowIndex = p[1]
			ans++
		}
	}
	return ans
}
