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
	var nextStart, lastNode *Node

	var makeConnect func(*Node)
	makeConnect = func(curNode *Node) {
		if curNode == nil {
			return
		}

		if nextStart == nil {
			nextStart = curNode
		}

		if lastNode != nil {
			lastNode.Next = curNode
		}

		lastNode = curNode
	}

	for start != nil {
		for p := start; p != nil; p = p.Next {
			makeConnect(p.Left)
			makeConnect(p.Right)
		}

		start = nextStart
		nextStart = nil
		lastNode = nil
	}

	return root
}
