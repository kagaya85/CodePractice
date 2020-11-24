package main

import "sort"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	for p := root; p != nil; p = p.Left {
		level++
	}

	// 使用二分搜索lelve层节点数量[2^(n-1), 2^n-1]
	return sort.Search(1<<level, func(i int) bool {
		if i <= 1<<(level-1) {
			return false
		}

		p := root
		bits := 1 << (level - 2)
		for p != nil && bits > 0 {
			if i&bits == 0 {
				p = p.Left
			} else {
				p = p.Right
			}
			bits >>= 1
		}
		return p == nil
	}) - 1
}
