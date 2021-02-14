/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	
	var ans [][]int
	
	var flag bool
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
		
		if flag {
			reverse(level)
		}
		ans = append(ans, append([]int{}, level...))
		
		flag = !flag
	}
	return ans
}

func reverse(level []int) {
	i, j := 0, len(level) - 1
	for i < j {
		level[i], level[j] = level[j], level[i]
		i++
		j--
	}
}