package main

import (
	"container/heap"
	"sort"
)

// KthLargest struct
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() (v interface{}) {
	a := kl.IntSlice
	v, kl.IntSlice = a[len(a)-1], a[:len(a)-1]
	return
}
