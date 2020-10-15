package main

// Node Definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	start := root
	var pre *Node
	for p := root; p.Left != nil; {
		if p != start {
			pre.Next = p.Left
		}
		p.Left.Next = p.Right

		if p.Next == nil {
			p = start.Left
			start = p
		} else {
			pre = p.Right
			p = p.Next
		}
	}

	return root
}
