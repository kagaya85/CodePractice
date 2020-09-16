package main

func main() {

}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	invertTree(root.Right)
	invertTree(root.Left)

	root.Left, root.Right = root.Right, root.Left

	return root
}
