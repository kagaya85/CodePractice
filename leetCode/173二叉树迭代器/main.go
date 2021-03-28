package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return BSTIterator{stack: stack}
}

func (b *BSTIterator) Next() int {
	p := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]
	for t := p.Right; t != nil; t = t.Left {
		b.stack = append(b.stack, t)
	}
	return p.Val
}

func (b *BSTIterator) HasNext() bool {
	return len(b.stack) > 0
}
