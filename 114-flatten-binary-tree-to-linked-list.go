/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten1(root *TreeNode) {
    if root == nil {
        return
    }

    for root != nil {
        if root.Left == nil {
            root = root.Right
            continue
        }

        pre := root.Left
        for pre.Right != nil {
            pre = pre.Right
        }

        pre.Right = root.Right
        root.Right = root.Left
        root.Left = nil
        root = root.Right
    }
}

func flatten(root *TreeNode) {
    if root == nil {
        return
    }

    var result []*TreeNode
    helper(root, &result)

    for i := 1; i < len(result); i++ {
        pre := result[i-1]
        pre.Left = nil
        pre.Right = result[i]
    }
}

func helper(root *TreeNode, nodes *[]*TreeNode) {
    if root == nil {
        return
    }
    *nodes = append(*nodes, root)
    helper(root.Left, nodes)
    helper(root.Right, nodes)
}