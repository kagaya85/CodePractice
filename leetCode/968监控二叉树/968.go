package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func minCameraCover(root *TreeNode) (ans int) {

	var lrd func(root *TreeNode) int
	lrd = func(root *TreeNode) int {
		if root == nil {
			return 1
		}

		left := lrd(root.Left)
		right := lrd(root.Right)

		if left == 0 || right == 0 {
			ans++
			return 2
		}

		if left == 1 && right == 1 {
			return 0
		}

		return 1
	}

	if lrd(root) == 0 {
		ans++
	}

	return
}
