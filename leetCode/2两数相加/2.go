package main

// ListNode def
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var flag bool

	if l1 == nil {
		l1, l2 = l2, l1
	}

	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		p1.Val += (p2.Val + boolToInt(flag))
		if p1.Val >= 10 {
			flag = true
			p1.Val -= 10
		} else {
			flag = false
		}

		if p1.Next == nil || p2.Next == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1.Next == nil && p2.Next == nil {
		if flag {
			p1.Next = &ListNode{Val: boolToInt(flag)}
		}
		return l1
	}

	if p1.Next == nil {
		p1.Next = p2.Next
	}
	p1.Next.Val += boolToInt(flag)

	for p1 = p1.Next; p1 != nil && p1.Val >= 10; p1 = p1.Next {
		p1.Val -= 10
		if p1.Next == nil {
			p1.Next = &ListNode{Val: 0}
		}
		p1.Next.Val++
	}

	return l1
}

func boolToInt(n bool) int {
	if n {
		return 1
	}
	return 0
}
