package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func pathSum(root *TreeNode, sum int) (ans [][]int) {

	tmpSum := 0
	tmpList := []int{}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		tmpSum += root.Val
		tmpList = append(tmpList, root.Val)

		if tmpSum == sum && root.Left == nil && root.Right == nil {
			ans = append(ans, append([]int(nil), tmpList...))
			tmpList = tmpList[:len(tmpList)-1]
			tmpSum -= root.Val
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		tmpList = tmpList[:len(tmpList)-1]
		tmpSum -= root.Val
	}

	dfs(root)

	return
}
