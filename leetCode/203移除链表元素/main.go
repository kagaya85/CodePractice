package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head = &ListNode{Next: head}
	p, q := head, head.Next

	for q != nil {
		if q.Val == val {
			p.Next = q.Next
		} else {
			p = p.Next
		}
		q = q.Next
	}

	return head.Next
}
