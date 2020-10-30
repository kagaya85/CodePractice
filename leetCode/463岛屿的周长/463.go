package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) (ans int) {
	width := len(grid[0])
	length := len(grid)

	for y, line := range grid {
		for x, unit := range line {
			if unit == 0 {
				continue
			}
			if x == 0 || line[x-1] == 0 {
				ans++
			}
			if x == width-1 || line[x+1] == 0 {
				ans++
			}
			if y == 0 || grid[y-1][x] == 0 {
				ans++
			}
			if y == length-1 || grid[y+1][x] == 0 {
				ans++
			}
		}
	}

	return
}
