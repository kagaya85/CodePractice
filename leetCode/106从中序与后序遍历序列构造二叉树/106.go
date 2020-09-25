package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	buildTree(inorder, postorder)

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	indexMap := map[int]int{}

	for i, num := range inorder {
		indexMap[num] = i
	}

	var build func(int, int) *TreeNode
	build = func(inorderLeft int, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}

		num := postorder[len(postorder)-1]
		index := indexMap[num]
		postorder = postorder[:len(postorder)-1]

		root := &TreeNode{Val: num}
		root.Right = build(index+1, inorderRight)
		root.Left = build(inorderLeft, index-1)
		return root
	}

	return build(0, len(inorder)-1)
}
