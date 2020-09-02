func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var     values1 i       values2 )
 int
nt
 (

    leafs(root1, &values1)
    leafs(root2, &values2)

    return values1 == values2
}

func leafs(root *TreeNode, result *int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *result = *result*10 + root.Val
        return
    }

    leafs(root.Left, result)
    leafs(root.Right, result)
}