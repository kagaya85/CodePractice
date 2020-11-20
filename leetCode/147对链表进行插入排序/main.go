package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func insertionSortList(head *ListNode) *ListNode {
	newHead := &ListNode{Next: nil}

	for head != nil {
		p := head
		head = head.Next
		q := newHead
		for q.Next != nil && q.Next.Val < p.Val {

			q = q.Next
		}
		p.Next = q.Next
		q.Next = p
	}

	return newHead.Next
}
