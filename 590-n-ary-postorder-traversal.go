package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	for _, node := range root.Children {
		result = append(result, postorder(node)...)
	}
	result = append(result, root.Val)
	return result
}

func postorder2(root *Node) []int {
	stack := NewStack()
	var result []int
	node := root

	stack.push(node)

	for !stack.isEmpty() {
		pop := stack.pop()
		if pop != nil {
			result = append(result, pop.Val)
			for i := 0; i < len(pop.Children); i++ {
				node := pop.Children[i]
				stack.push(node)
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

type Stack struct {
	values []*Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) push(v *Node) {
	s.values = append(s.values, v)
}

func (s *Stack) pop() *Node {
	if len(s.values) == 0 {
		return nil
	}
	x, v := s.values[len(s.values)-1], s.values[:len(s.values)-1]
	s.values = v
	return x
}
