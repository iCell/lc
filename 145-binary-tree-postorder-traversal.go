package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root != nil {
		result = append(result, preorderTraversal(root.Left)...)
		result = append(result, preorderTraversal(root.Right)...)
		result = append(result, root.Val)
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	stack := NewStack()
	node := root

	for !stack.isEmpty() || node != nil {
		for node != nil {
			result = append(result, node.Val)
			stack.push(node)
			node = node.Right
		}
		pop := stack.pop()
		node = pop.Left
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
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
