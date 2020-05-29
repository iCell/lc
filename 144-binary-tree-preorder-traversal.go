package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, root.Val)
		result = append(result, preorderTraversal(root.Left)...)
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}

func preorderTraversal2(root *TreeNode) []int {
	var result []int
	stack := NewStack()
	node := root
	for stack.isEmpty() || node != nil {
		for node != nil {
			result = append(result, node.Val)
			stack.push(node)
			node = node.Left
		}
		top := stack.pop()
		node = top.Right
	}
	return result
}

type Stack struct {
	values []*TreeNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) push(v *TreeNode) {
	s.values = append(s.values, v)
}

func (s *Stack) pop() *TreeNode {
	if len(s.values) == 0 {
		return nil
	}
	x, v := s.values[len(s.values)-1], s.values[:len(s.values)-1]
	s.values = v
	return x
}
