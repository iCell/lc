/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
	var sum int
	helper(root, &sum)
	return sum
}

func helper(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	
	left, right := helper(root.Left, sum), helper(root.Right, sum)
	tilt := abs(left, right)
	*sum = *sum + tilt
	return left + right + root.Val
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}