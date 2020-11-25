/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	return helper(root, memo)
}

func helper(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	
	cached, exist := memo[root]
	if exist {
		return cached
	}
	
	i := root.Val
	if root.Left != nil {
		i += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		i += rob(root.Right.Left) + rob(root.Right.Right)
	}
	
	j := rob(root.Left) + rob(root.Right)
	ans := max(i, j)
	
	memo[root] = ans
	
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}