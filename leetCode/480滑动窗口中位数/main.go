package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
	size int
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) {
	a := h.IntSlice
	v, h.IntSlice = a[len(a)-1], a[:len(a)-1]
	return
}
func (h *hp) push(v int) { h.size++; heap.Push(h, v) }
func (h *hp) pop() int   { h.size--; return heap.Pop(h).(int) }

func medianSlidingWindow(nums []int, k int) []float64 {
	delayed := map[int]int{}
	small, large := &hp{}, &hp{} // small是绝对值的大根堆，插入元素为负数来模拟

	// 删除堆顶可以删除的元素
	prune := func(h *hp) {
		for h.Len() > 0 {
			num := h.IntSlice[0] // 获取堆顶元素
			if h == small {      // 用负数模拟大根碓
				num = -num
			}
			if d, ok := delayed[num]; ok {
				if d > 1 {
					delayed[num]--
				} else {
					delete(delayed, num)
				}
				heap.Pop(h)
			} else {
				break
			}
		}
	}

	// 每次插入一个元素后，调整两个堆的元素个数
	makeBalance := func() {
		if small.size > large.size+1 { // small比large多两个
			large.push(-small.pop())
			prune(small)
		} else if small.size < large.size { // small比large少一个
			small.push(-large.pop())
			prune(large)
		}
	}

	// 插入元素
	insert := func(num int) {
		if small.size == 0 || num <= -small.IntSlice[0] {
			small.push(-num)
		} else {
			large.push(num)
		}
		makeBalance()
	}

	// 删除元素
	remove := func(num int) {
		delayed[num]++
		if num <= -small.IntSlice[0] {
			small.size--
			if num == -small.IntSlice[0] { // 如果就是堆顶元素则立即删除
				prune(small)
			}
		} else {
			large.size--
			if num == large.IntSlice[0] {
				prune(large)
			}
		}
		makeBalance()
	}

	// 获取中位数
	getMedian := func() float64 {
		if k&1 > 0 { // 窗口长度为奇数
			return float64(-small.IntSlice[0])
		}
		return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
	}

	for _, num := range nums[:k] { // 先插入k个元素
		insert(num)
	}

	n := len(nums)
	ans := make([]float64, 1, n-k+1)
	ans[0] = getMedian()
	for i := k; i < n; i++ {
		insert(nums[i])
		remove(nums[i-k])
		ans = append(ans, getMedian())
	}

	return ans
}
