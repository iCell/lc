func maxAncestorDiff(root *TreeNode) int {
	var diff int
	helper(root, root.Val, root.Val, &diff)
	return diff
}

func helper(root *TreeNode, min, max int, diff *int) {
	if root == nil {
		return
	}
	if root.Val < min {
		min = root.Val
	}
	if root.Val > max {
		max = root.Val
	}
	if max - min > *diff {
		*diff = max - min
	}
	helper(root.Left, min, max, diff)
	helper(root.Right, min, max, diff)
}