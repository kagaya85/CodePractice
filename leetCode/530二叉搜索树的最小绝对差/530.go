package main

import "math"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt32
	last := -1

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if last != -1 {
			if diff := root.Val - last; diff < min {
				min = diff
			}
		}
		last = root.Val
		dfs(root.Right)
	}

	dfs(root)

	return min
}
