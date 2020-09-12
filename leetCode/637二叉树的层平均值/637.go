package main

import (
	"fmt"
	"container/list"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	
}

func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
		
		nextlevel := list.New()
		nextlevel.PushBack(root)

		for nextlevel.Len() > 0 {
			curlevel := nextlevel
			nextlevel = list.New()
			sum := 0
			len := curlevel.Len()
			
			for e := curlevel.Front(); e != nil; e = e.Next() {
				p := e.Value.(*TreeNode)
				sum += p.Val
				if p.Left != nil {
					nextlevel.PushBack(p.Left)
				}
				if p.Right != nil {
					nextlevel.PushBack(p.Right)
				}
			}
			
			ans = append(ans, float64(sum)/float64(len))
			
		}
		
		return ans
}