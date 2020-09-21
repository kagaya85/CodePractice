package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		sum += root.Val
		root.Val = sum
		dfs(root.Left)
	}

	dfs(root)

	return root
}
