package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Next: head}

	p, q := newHead, newHead.Next
	for q != nil {
		r := q.Next
		for r != nil && q.Val == r.Val {
			q, r = r, r.Next
		}
		if p.Next != q {
			p.Next = r
			q = r
		} else {
			p, q = q, r
		}
	}

	return newHead.Next
}
