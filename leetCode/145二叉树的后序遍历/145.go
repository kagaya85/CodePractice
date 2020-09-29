package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func postorderTraversal(root *TreeNode) (ans []int) {

	stack := []*TreeNode(nil)
	p := root
	var prev *TreeNode

	for p != nil || len(stack) != 0 {

		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		// pop
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if p.Right == nil || p.Right == prev {
			// visit
			ans = append(ans, p.Val)
			prev = p
			p = nil
		} else {
			stack = append(stack, p)
			p = p.Right
		}
	}

	return
}
