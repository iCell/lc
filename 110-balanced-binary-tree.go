func isBalanced(root *TreeNode) bool {
    return helper(root) >= 0
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }

    l, r := helper(root.Left), helper(root.Right)

    if l == -1 || r == -1 || abs(l, r) > 1 {
        return -1
    }

    return max(l, r) + 1
}

func abs(x, y int) int {
    if x-y < 0 {
        return y - x
    }
    return x - y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}