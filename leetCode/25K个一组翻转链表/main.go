package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	prev := &ListNode{Next: head}
	head, tail := prev, prev
	for {
		for i := 0; i < k; i++ {
			if tail.Next == nil {
				return prev.Next
			}
			tail = tail.Next
		}
		head.Next, tail = reverseGroup(head.Next, tail)
		head = tail
	}
}

func reverseGroup(head, tail *ListNode) (*ListNode, *ListNode) {
	prev, p := tail.Next, head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}
