/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	nodes := helper(root)
	if len(nodes) == 0 {
		return nil
	}

	for i := 0; i < len(nodes) - 1; i++ {
		cur, next := nodes[i], nodes[i+1]
		cur.Left, cur.Right = nil, next
		next.Left, next.Right = nil, nil
	}
	return nodes[0]
}

func helper(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var result []*TreeNode
	result = append(result, helper(root.Left)...)
	result = append(result, root)
	result = append(result, helper(root.Right)...)
	return result
}