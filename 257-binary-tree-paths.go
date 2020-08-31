/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }
    paths := []string{}
    helper(root, fmt.Sprintf("%d", root.Val), &paths)
    return paths
}

func helper(root *TreeNode, path string, paths *[]string) {
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil {
        *paths = append(*paths, path)
        return
    }
    
    if root.Left != nil {
        newPath := fmt.Sprintf("%s->%d", path, root.Left.Val)
        helper(root.Left, newPath, paths)
    }
    
    if root.Right != nil {
        newPath := fmt.Sprintf("%s->%d", path, root.Right.Val)
        helper(root.Right, newPath, paths)
    }
}