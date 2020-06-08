package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	queue := &Queue{}

	queue.InQueue(root)
	for !queue.IsEmpty() {
		level := queue.Size()
		var max *int
		for i := 0; i < level; i++ {
			node := queue.DeQueue()
			if node.Left != nil {
				queue.InQueue(node.Left)
			}
			if node.Right != nil {
				queue.InQueue(node.Right)
			}
			if max == nil {
				max = &node.Val
			} else if *max < node.Val {
				max = &node.Val
			}
		}
		result = append(result, *max)
	}

	return result
}

func largestValues(root *TreeNode) []int {
	var result []int
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

		if len(nodes) == 0 {
			continue
		}

		max := nodes[0].Val
		for i, node := range nodes {
			if i > 0 && node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue.InQueue(node.Left)
			}
			if node.Right != nil {
				queue.InQueue(node.Right)
			}
		}
		result = append(result, max)
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
