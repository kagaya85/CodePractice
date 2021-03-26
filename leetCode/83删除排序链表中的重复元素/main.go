package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for p, q := head, head.Next; q != nil; q = q.Next {
		if p.Val == q.Val {
			p.Next = q.Next
		} else {
			p = q
		}
	}
	return head
}
