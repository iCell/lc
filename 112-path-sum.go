
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	left := hasPathSum(root.Left, sum-root.Val)
	right := hasPathSum(root.Right, sum-root.Val)
	return left || right
}
