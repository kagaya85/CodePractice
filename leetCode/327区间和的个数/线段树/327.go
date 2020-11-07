package main

import "math"

func main() {

}

type node struct {
	l, r, val int
	lo, ro    *node
}

func (o *node) insert(val int) {
	o.val++
	if o.l == o.r {
		return
	}
	m := (o.l + o.r) >> 1
	if val <= m {
		if o.lo == nil {
			o.lo = &node{l: o.l, r: m}
		}
		o.lo.insert(val)
	} else {
		if o.ro == nil {
			o.ro = &node{l: m + 1, r: o.r}
		}
		o.ro.insert(val)
	}
}

func (o *node) query(l, r int) (res int) {
	if o == nil || l > o.r || r < o.l {
		return
	}
	if l <= o.l && o.r <= r {
		return o.val
	}
	return o.lo.query(l, r) + o.ro.query(l, r)
}

func countRangeSum(nums []int, lower, upper int) (cnt int) {
	preSum := make([]int, len(nums)+1)
	for i, v := range nums {
		preSum[i+1] = preSum[i] + v
	}

	lbound, rbound := math.MaxInt64, -math.MaxInt64
	for _, sum := range preSum {
		lbound = min(lbound, sum, sum-lower, sum-upper)
		rbound = max(rbound, sum, sum-lower, sum-upper)
	}

	root := &node{l: lbound, r: rbound}
	for _, sum := range preSum {
		left, right := sum-upper, sum-lower
		cnt += root.query(left, right)
		root.insert(sum)
	}
	return
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
