package main

import (
	"sort"
)

type Stack struct {
	sort.IntSlice
}

func (s *Stack) Push(x int) {
	s.IntSlice = append(s.IntSlice, x)
}

func (s *Stack) Pop() int {
	x := s.IntSlice[s.Len()-1]
	s.IntSlice = s.IntSlice[:s.Len()-1]
	return x
}

func (s *Stack) Peek() int {
	return s.IntSlice[s.Len()-1]
}

type MyQueue struct {
	in, out Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Stack{}, Stack{}}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.out.Len() == 0 {
		for q.in.Len() != 0 {
			q.out.Push(q.in.Pop())
		}
	}
	if q.out.Len() == 0 {
		return 0
	}
	return q.out.Pop()
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.out.Len() == 0 {
		for q.in.Len() != 0 {
			q.out.Push(q.in.Pop())
		}
	}
	if q.out.Len() == 0 {
		return 0
	}
	return q.out.Peek()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.in.Len() == 0 && q.out.Len() == 0
}
