package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

func inorderTraversal2(root *TreeNode) []int {
	var result []int

	node := root
	stack := NewStack()
	for len(stack.values) != 0 || node != nil {
		for node != nil {
			stack.push(node)
			node = node.Left
		}
		top := stack.pop()
		result = append(result, top.Val)
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
