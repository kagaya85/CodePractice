package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	n := len(h.IntSlice)
	v := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return v
}

func minRefuelStops(target int, startFuel int, stations [][]int) (ans int) {
	fuel, prev, h := startFuel, 0, &hp{}
	n := len(stations)
	for i := 0; i <= n; i++ {
		curr := target
		if i < n {
			curr = stations[i][0]
		}

		fuel -= curr - prev
		for fuel < 0 && h.Len() > 0 {
			fuel += heap.Pop(h).(int)
			ans++
		}

		if fuel < 0 {
			return -1
		}

		if i < n {
			heap.Push(h, stations[i][1])
			prev = curr
		}
	}
	return
}

func main() {

}
