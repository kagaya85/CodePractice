package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	for p := root; p != nil; {
		if p.Val < val {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			}
			p = p.Right
		} else if p.Val > val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			}
			p = p.Left
		} else {
			break
		}
	}

	return root
}
