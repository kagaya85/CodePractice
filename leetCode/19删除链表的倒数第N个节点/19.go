package main

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headNode := &ListNode{0, head}
	p, q := headNode, head

	for i := 0; i < n; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	// delete p next
	p.Next = p.Next.Next

	return headNode.Next
}
