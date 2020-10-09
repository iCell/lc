/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    paths := make([][]int, 0)
    helper(root, sum, []int{}, &paths)
    return paths
}

func helper(root *TreeNode, sum int, path []int, paths *[][]int) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        if sum == root.Val {
            path = append(path, root.Val)
            *paths = append(*paths, append([]int{}, path...))
        }
        return
    }

    path = append(path, root.Val)
    if root.Left != nil {
        helper(root.Left, sum-root.Val, path, paths)
    }
    if root.Right != nil {
        helper(root.Right, sum-root.Val, path, paths)
    }
}