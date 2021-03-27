package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	var n int
	p := head
	for {
		n++
		if p.Next == nil {
			p.Next = head
			break
		}
		p = p.Next
	}

	l := n - k%n
	for i := 0; i < l; i++ {
		p = p.Next
	}
	head = p.Next
	p.Next = nil
	return head
}
