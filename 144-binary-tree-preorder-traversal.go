package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	helper(&result, root)
	return result
}

func helper(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	helper(result, root.Left)
	helper(result, root.Right)
}

func preorderTraversal2(root *TreeNode) []int {
	var result []int
	stack := NewStack()
	node := root
	for stack.isEmpty() || node != nil {

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
