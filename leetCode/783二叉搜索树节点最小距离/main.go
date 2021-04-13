package main

import "math"

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	pre, ans := -1, math.MaxInt64
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			ans = min(ans, node.Val-pre)
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
