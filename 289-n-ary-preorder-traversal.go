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
