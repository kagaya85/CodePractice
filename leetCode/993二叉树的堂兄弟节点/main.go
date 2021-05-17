package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node, parent *TreeNode, depth int) bool
	dfs = func(node, parent *TreeNode, depth int) bool {
		if node == nil {
			return false
		}
		if node.Val == x {
			xFound, xParent, xDepth = true, parent, depth
		} else if node.Val == y {
			yFound, yParent, yDepth = true, parent, depth
		}
		if xFound && yFound {
			return true
		}
		if dfs(node.Left, node, depth+1) || dfs(node.Right, node, depth+1) {
			return true
		}
		return false
	}
	dfs(root, nil, 0)

	return xDepth == yDepth && xParent != yParent
}
