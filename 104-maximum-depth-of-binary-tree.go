package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	left := helper(root.Left, depth+1)
	right := helper(root.Right, depth+1)
	return int(math.Max(float64(left), float64(right)))
}
