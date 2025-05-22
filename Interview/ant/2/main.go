// 层序遍历打印 二叉树
package main

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func LevelOrderTraversal(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level[i] = node.Value

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
