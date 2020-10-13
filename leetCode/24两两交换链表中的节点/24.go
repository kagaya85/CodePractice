package main

// ListNode Defination
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	p := &ListNode{Next: head}
	q := head
	head = q.Next

	for q != nil && q.Next != nil {
		p.Next = q.Next
		q.Next = q.Next.Next
		p.Next.Next = q

		p = q
		q = q.Next
	}

	return head
}
