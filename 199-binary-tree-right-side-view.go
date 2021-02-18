/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	queue, result := []*TreeNode{root}, make([]int, 0)
	
	for len(queue) != 0 {
		size, level := len(queue), make([]int, 0)
		for size > 0 {
			first, temp := queue[0], queue[1:]
			queue = temp
			level = append(level, first.Val)
		
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			size--
		}
		if len(level) > 0 {
			result = append(result, level[len(level)-1])
		}
	}
	
	return result
}