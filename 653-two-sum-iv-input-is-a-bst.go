func findTarget(root *TreeNode, k int) bool {
    memo := make(map[int]bool)
    return helper(root, k, memo)
}

func helper(root *TreeNode, k int, memo map[int]bool) bool {
    if root == nil {
        return false
    }
    if memo[k-root.Val] == true {
        return true
    }
    memo[root.Val] = true
    return helper(root.Left, k, memo) || helper(root.Right, k, memo)
}