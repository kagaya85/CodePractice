package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

// 最大堆
func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func lastStoneWeight(stones []int) int {
	h := &hp{stones}
	heap.Init(h)

	for h.Len() > 1 {
		x, y := h.pop(), h.pop()
		if x > y {
			h.push(x - y)
		}
	}

	if h.Len() > 0 {
		return h.IntSlice[0]
	}
	return 0
}
