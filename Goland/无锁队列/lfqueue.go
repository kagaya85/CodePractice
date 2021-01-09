package lfqueue

// lock-free queue

import (
	"sync/atomic"
	"unsafe"
)

type LFQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func NewLFQueue() *LFQueue {
	h := unsafe.Pointer(&node{})
	return &LFQueue{head: h, tail: h}
}

func (q *LFQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) { // 检查tail是否被修改
			if next == nil { // 检查tail的next是否还为nil
				if cas(&tail.next, next, n) { // 尝试修改tail的next指向新加入的节点
					cas(&q.tail, tail, n) // 添加新节点成功后修改tail指向新的队列尾节点（若不成功，说明tail节点已被更新）
					return
				}
			} else {
				cas(&q.tail, tail, next) // 若next不为nil，代表tail节点已经滞后，尝试更新tail
			}
		}
	}
}

func (q *LFQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // 检查head是否被修改
			if head == tail { // 队列为空或者tail还未来得及修改
				if next == nil { // 队列为空
					return nil
				}
				cas(&q.tail, tail, next) // 更新tail
			} else {
				// 先读取要出队列的数值
				v := next.value
				if cas(&q.head, head, next) { // 尝试更新头节点，确保该元素还在队列头并弹出
					return v // 该元素还在队列头，弹出并返回该值
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
