package main

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := findMiddleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	head = mergeList(l1, l2)
}

func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func mergeList(L1, L2 *ListNode) (head *ListNode) {
	if L1 == nil {
		L1, L2 = L2, L1
	}

	head = L1

	for L1 != nil && L2 != nil {
		tmp1 := L1.Next
		tmp2 := L2.Next

		L1.Next = L2
		L2.Next = tmp1

		L1 = tmp1
		L2 = tmp2
	}

	return
}
