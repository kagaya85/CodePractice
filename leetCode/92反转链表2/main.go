package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	n := right - left
	if n < 1 {
		return head
	}
	newHead := &ListNode{Next: head}

	p := newHead
	for i := 0; i < left-1; i++ {
		p = p.Next
	}

	front := p
	p = p.Next
	q, r := p.Next, p.Next.Next
	for i := 0; i < n && p != nil && q != nil; i++ {
		q.Next = p
		p, q = q, r
		if r != nil {
			r = r.Next
		}
	}

	front.Next.Next = q
	front.Next = p

	return newHead.Next
}
