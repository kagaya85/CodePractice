package main

// TreeNode define for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, val int) {
		if root.Left == nil && root.Right == nil {
			ans += val
			return
		}
		if root.Left != nil {
			dfs(root.Left, val*10+root.Left.Val)
		}
		if root.Right != nil {
			dfs(root.Right, val*10+root.Right.Val)
		}
	}

	if root != nil {
		dfs(root, root.Val)
	}

	return
}
