package main

import (
	"math"
	"math/rand"
)

type node struct {
	child    [2]*node
	priority int
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.child[d^1]
	o.child[d^1] = x.child[d]
	x.child[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val}
	}
	if d := o.cmp(val); d >= 0 {
		o.child[d] = t._put(o.child[d], val)
		if o.child[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.child[0]
		case c > 0:
			o = o.child[1]
		default:
			return o
		}
	}
	return
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	ans := math.MinInt64
	for i := range matrix { // 枚举上边界
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				sum[c] += v // 更新每列的元素和
			}
			t := &treap{}
			t.put(0)
			s := 0
			for _, v := range sum {
				s += v
				if lb := t.lowerBound(s - k); lb != nil {
					ans = max(ans, s-lb.val)
				}
				t.put(s)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
