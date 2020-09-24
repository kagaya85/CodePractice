package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func findMode(root *TreeNode) []int {

	maxCount := 0
	count := 0
	num := 0
	var modeNum []int

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)

		if num == root.Val {
			count++
		} else {
			num, count = root.Val, 1
		}

		if count == maxCount {
			modeNum = append(modeNum, root.Val)
		}

		if count > maxCount {
			modeNum, maxCount = []int{root.Val}, count
		}

		inorder(root.Right)
	}

	inorder(root)

	return modeNum
}
