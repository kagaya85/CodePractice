package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{node.Val}
		}
		return append(dfs(node.Left), dfs(node.Right)...)
	}

	leafs1, leafs2 := dfs(root1), dfs(root2)
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i, v := range leafs1 {
		if v != leafs2[i] {
			return false
		}
	}
	return true
}
