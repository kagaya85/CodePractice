package main

import "container/list"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	stack := list.New()
	p := root

	var ans []int

	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		if stack.Len() != 0 {
			e := stack.Back()
			p = e.Value.(*TreeNode)
			ans = append(ans, p.Val)
			stack.Remove(e)
			p = p.Right
		}
	}

	return ans
}
