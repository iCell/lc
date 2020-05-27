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
