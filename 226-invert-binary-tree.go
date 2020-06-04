package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree2(root.Left)
	root.Left, root.Right = root.Right, root.Left
	invertTree2(root.Left)

	return root
}

func invertTree3(root *TreeNode) *TreeNode {
	node := root
	stack := NewStack()
	for len(stack.values) != 0 || node != nil {
		for node != nil {
			stack.push(node)
			node = node.Left
		}
		top := stack.pop()

		top.Left, top.Right = top.Right, top.Left
		node = top.Left
	}

	return root
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
