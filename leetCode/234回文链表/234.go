package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome(head *ListNode) bool {
	list := make([]int, 0, 100)
	n := 0

	for head != nil {
		list = append(list, head.Val)
		n++
		head = head.Next
	}

	var left, right int
	if n%2 == 0 {
		right = n / 2
		left = right - 1
	} else {
		right = n/2 + 1
		left = right - 2
	}

	for left >= 0 && right < n {
		if list[left] != list[right] {
			return false
		}
		left--
		right++
	}

	return true
}
