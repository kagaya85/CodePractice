package dijkstra

import (
	"container/heap"
	"math"
)

type point struct {
	x, y, maxDiff int
}

type hp []point

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].maxDiff < h[j].maxDiff }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(point)) }
func (h *hp) Pop() (v interface{}) {
	a := *h
	*h, v = a[:len(a)-1], a[len(a)-1]
	return
}

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	maxDiff := make([][]int, rows)
	for i := range maxDiff {
		maxDiff[i] = make([]int, cols)
		for j := range maxDiff[i] {
			maxDiff[i][j] = math.MaxInt64
		}
	}

	maxDiff[0][0] = 0
	h := &hp{{}}
	for {
		p := heap.Pop(h).(point)
		if p.x == rows-1 && p.y == cols-1 {
			return p.maxDiff
		}
		if maxDiff[p.x][p.y] < p.maxDiff {
			continue
		}
		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < rows && 0 <= y && y < cols {
				if diff := max(p.maxDiff, abs(heights[x][y]-heights[p.x][p.y])); diff < maxDiff[x][y] {
					maxDiff[x][y] = diff
					heap.Push(h, point{x, y, diff})
				}
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
