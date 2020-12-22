package main

// TreeNode Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		nodes := queue
		queue = nil
		for _, n := range nodes {
			vals = append(vals, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		if level%2 != 0 {
			for i, l := 0, len(vals); i < l/2; i++ {
				vals[i], vals[l-i-1] = vals[l-i-1], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}
