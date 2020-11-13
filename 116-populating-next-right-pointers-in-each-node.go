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
		var level []*Node
		
		size := len(queue)
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
			
			size--
		}
		
		for i := 1; i < len(level); i++ {
			level[i-1].Next = level[i]
		}
	}
	
	return root
}

