package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		left := minDepth(root.Left)
		right := minDepth(root.Right)
		return int(math.Min(float64(left), float64(right))) + 1
	} else {
		left := minDepth(root.Left)
		right := minDepth(root.Right)
		return int(math.Max(float64(left), float64(right))) + 1
	}
}
