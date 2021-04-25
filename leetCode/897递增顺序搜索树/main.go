package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) (head *TreeNode) {
	queue := []*TreeNode{root}
	q := root.Left
	for q != nil {
		queue = append(queue, q)
		q = q.Left
	}
	var p *TreeNode
	for len(queue) > 0 {
		q := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if head == nil {
			head, p = q, q
		} else {
			p.Right = q
			p = p.Right
			q.Left = nil
		}
		if q.Right != nil {
			q = q.Right
			for q != nil {
				queue = append(queue, q)
				q = q.Left
			}
		}
	}
	return
}
