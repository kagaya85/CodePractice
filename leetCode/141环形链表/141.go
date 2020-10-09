package main

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 龟兔赛跑算法
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
