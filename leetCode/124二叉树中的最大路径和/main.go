/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))

		res = max(res, node.Val+left+right)

		return node.Val + max(left, right)
	}
	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}