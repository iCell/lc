/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	
	queue := make([]*Node, 0)
	queue = append(queue, root)
	
	for len(queue) != 0 {
		size, level := len(queue), make([]*Node, 0)
		for size > 0 {
			first, temp := queue[0], queue[1:]
			queue = temp
			level = append(level, first)
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			size -= 1
		}
		for i := 0; i < len(level) - 1; i++ {
			level[i].Next = level[i+1]
		}
	}
	
	return root
}