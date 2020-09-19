package main

func main() {

}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}

	if root.Left != nil && isLeft(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func isLeft(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
