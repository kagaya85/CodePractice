package main

import "container/heap"

type block struct {
	x, y, height int
}

type hp []block

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height < h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(block)) }
func (h *hp) Pop() (v interface{}) {
	a := *h
	*h, v = a[:len(a)-1], a[len(a)-1]
	return
}

func swimInWater(grid [][]int) (ans int) {

	rows, cols := len(grid), len(grid[0])

	arrived := make([][]bool, rows)
	for i := range arrived {
		arrived[i] = make([]bool, cols)
	}
	h := &hp{{0, 0, grid[0][0]}}
	arrived[0][0] = true
	direction := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for {
		blk := heap.Pop(h).(block)
		ans = max(ans, blk.height)
		if blk.x == cols-1 && blk.y == rows-1 {
			return
		}
		for _, d := range direction {
			if x, y := blk.x+d[0], blk.y+d[1]; x >= 0 && x < cols && y >= 0 && y < rows && !arrived[y][x] {
				arrived[y][x] = true
				heap.Push(h, block{
					x:      x,
					y:      y,
					height: grid[y][x],
				})
			}
		}
	}
}

func max(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
