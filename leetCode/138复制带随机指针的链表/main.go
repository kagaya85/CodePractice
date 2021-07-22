package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := map[*Node]*Node{}
	for p := head; p != nil; p = p.Next {
		nodeMap[p] = new(Node)
	}
	for p := head; p != nil; p = p.Next {
		q := nodeMap[p]
		q.Val = p.Val
		q.Next = nodeMap[p.Next]
		q.Random = nodeMap[p.Random]
	}
	return nodeMap[head]
}
