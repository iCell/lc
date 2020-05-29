package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	for _, node := range root.Children {
		result = append(result, preorder(node)...)
	}
	return result
}

func preorder(root *Node) []int {
	stack := NewStack()
	var result []int
	node := root

	stack.push(node)

	for !stack.isEmpty() {
		pop := stack.pop()
		if pop != nil {
			result = append(result, pop.Val)
			for i := len(pop.Children) - 1; i >= 0; i-- {
				node := pop.Children[i]
				stack.push(node)
			}
		}
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
