package main

type NestedInteger interface {
	IsInteger() bool
	GetInteger() int
	SetInteger(value int)
	Add(elem NestedInteger)
	GetList() []*NestedInteger
}

type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (n *NestedIterator) Next() int {
	top := n.stack[len(n.stack)-1]
	if len(top) > 1 {
		n.stack[len(n.stack)-1] = top[1:]
	} else {
		n.stack = n.stack[:len(n.stack)-1]
	}
	return top[0].GetInteger()
}

func (n *NestedIterator) HasNext() bool {
	for len(n.stack) > 0 {
		top := n.stack[len(n.stack)-1]
		if len(top) == 0 {
			n.stack = n.stack[:len(n.stack)-1]
			continue
		}

		next := top[0]
		if next.IsInteger() {
			return true
		}
		n.stack[len(n.stack)-1] = top[1:]
		n.stack = append(n.stack, next.GetList())
	}
	return false
}
