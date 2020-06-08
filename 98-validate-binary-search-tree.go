package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(nil, nil, root)
}

func helper(min, max *TreeNode, node *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}

	if !helper(min, node, node.Left) {
		return false
	}
	if !helper(node, max, node.Right) {
		return false
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	stack := NewStack()
	node := root
	var last *TreeNode
	for node != nil || !stack.isEmpty() {
		for node != nil {
			stack.push(node)
			node = node.Left
		}
		pop := stack.pop()

		if last != nil && pop != nil && last.Val >= pop.Val {
			return false
		}
		last = pop
		node = pop.Right
	}
	return true
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

// stupid solution
func isValidBST2(root *TreeNode) bool {
	values := make([]int, 0)
	reverse(root, &values)

	for i, num := range values {
		if i > 0 && num <= values[i-1] {
			return false
		}
	}
	return true
}

func reverse(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}

	reverse(root.Left, values)
	*values = append(*values, root.Val)
	reverse(root.Right, values)
}
