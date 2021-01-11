package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap struct {
	sort.IntSlice
}

// Push elem
func (h *MinHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

// Pop elem
func (h *MinHeap) Pop() interface{} {
	x := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return x
}

func smallestStringWithSwaps(s string, pairs [][]int) string {

	n := len(s)
	pMap := map[int]*MinHeap{}
	parents := make([]int, n)

	for i := range parents {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if x != parents[x] {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	union := func(from, to int) {
		parents[find(from)] = find(to)
	}

	for _, p := range pairs {
		union(p[0], p[1])
		fmt.Println(parents)
	}

	for i := range parents {
		p := find(i)
		mh, has := pMap[p]
		if !has {
			mh = new(MinHeap)
			pMap[p] = mh
		}
		heap.Push(mh, int(s[i]))
	}

	ans := make([]byte, len(s))
	for i := range ans {
		ans[i] = byte(heap.Pop(pMap[parents[i]]).(int))
	}

	return string(ans)
}
