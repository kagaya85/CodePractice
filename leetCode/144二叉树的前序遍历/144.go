package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func preorderTraversal(root *TreeNode) (ans []int) {
	var stack []*TreeNode

	stack = append(stack, root)

	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for p != nil {
			ans = append(ans, p.Val)
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			p = p.Left
		}
	}

	return
}
