package main

// ListNode Definition
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	leftHead, rightHead := &ListNode{}, &ListNode{}
	left, right := leftHead, rightHead
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			left.Next = p
			left = left.Next
		} else {
			right.Next = p
			right = right.Next
		}
	}

	left.Next = rightHead.Next
	right.Next = nil

	return leftHead.Next
}
