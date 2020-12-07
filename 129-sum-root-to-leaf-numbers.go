/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var sum int
	helper(root, 0, &sum)
	return sum
}

func helper(root *TreeNode, temp int, sum *int) {
	if root == nil {
		return
	}
	temp = temp * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum = *sum + temp
		return
	}
	helper(root.Left, temp, sum)
	helper(root.Right, temp, sum)
}