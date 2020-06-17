package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node1 := &TreeNode{Val: 1}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}

	node3.Left = node5
	node3.Right = node1
	node5.Left = node6
	node5.Right = node2
	node2.Left = node7
	node2.Right = node4
	node1.Left = node0
	node1.Right = node8

	fmt.Println(lowestCommonAncestor(node3, node5, node4).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if root == nil && left == nil {
		return nil
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
