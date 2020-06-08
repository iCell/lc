package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := &Queue{}

	queue.InQueue(root)
	for !queue.IsEmpty() {
		level := queue.Size()
		values := make([]int, 0, level)
		for i := 0; i < level; i++ {
			node := queue.DeQueue()
			if node.Left != nil {
				queue.InQueue(node.Left)
			}
			if node.Right != nil {
				queue.InQueue(node.Right)
			}
			values = append(values, node.Val)
		}
		result = append(result, values)
	}

	return result
}

func levelOrder2(root *TreeNode) [][]int {
	var result [][]int
	queue := &Queue{}

	queue.InQueue(root)
	for !queue.IsEmpty() {
		nodes := make([]*TreeNode, 0, queue.Size())
		for {
			node := queue.DeQueue()
			if node == nil {
				break
			}
			nodes = append(nodes, node)
		}

		values := make([]int, 0, queue.Size())
		for _, node := range nodes {
			values = append(values, node.Val)
			if node.Left != nil {
				queue.InQueue(node.Left)
			}
			if node.Right != nil {
				queue.InQueue(node.Right)
			}
		}
		if len(values) != 0 {
			result = append(result, values)
		}
	}

	return result
}

type Queue struct {
	values []*TreeNode
}

func (q *Queue) InQueue(n *TreeNode) {
	q.values = append(q.values, n)
}

func (q *Queue) DeQueue() *TreeNode {
	if q.IsEmpty() {
		return nil
	}
	first, values := q.values[0], q.values[1:]
	q.values = values
	return first
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue) Size() int {
	return len(q.values)
}
