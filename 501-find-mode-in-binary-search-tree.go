/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	memo := make(map[int]int)
	helper(root, memo)
	var max int
	for _, cnt := range memo {
		if cnt > max {
			max = cnt
		}
	}
	var ans []int
	for node, cnt := range memo {
		if cnt == max {
			ans = append(ans, node)
		}
	}
	return ans
}

func helper(root *TreeNode, visited map[int]int) {
	if root == nil {
		return
	}
	
	helper(root.Left, visited)
	visited[root.Val] += 1
	helper(root.Right, visited)
}