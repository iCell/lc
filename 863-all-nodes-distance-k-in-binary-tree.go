/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil {
		return nil
	}
	parents, visited := make(map[*TreeNode]*TreeNode), make(map[*TreeNode]bool)
	findParents(root, parents)
	
	ans := make([]int, 0)
	search(target, K, parents, visited, &ans)
	return ans
}

func findParents(node *TreeNode, parents map[*TreeNode]*TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		parents[node.Left] = node
	}
	if node.Right != nil {
		parents[node.Right] = node
	}
	findParents(node.Left, parents)
	findParents(node.Right, parents)
}

func search(
	node *TreeNode, k int, 
	parents map[*TreeNode]*TreeNode, visited map[*TreeNode]bool, 
	result *[]int,
) {
	if node == nil || k < 0 || visited[node] {
		return
	}
	visited[node] = true
	if k == 0  {
		*result = append(*result, node.Val)
		return
	}
	search(node.Left, k - 1, parents, visited, result)
	search(parents[node], k - 1, parents, visited, result)
	search(node.Right, k - 1, parents, visited, result)
}