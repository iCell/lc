package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	current := root
	carry := 0

	node1, node2 := l1, l2
	for node1 != nil && node2 != nil {
		val := node1.Val + node2.Val + carry
		if val >= 10 {
			val = val % 10
			carry = 1
		} else {
			carry = 0
		}
		node := &ListNode{Val: val}

		current.Next = node
		current = current.Next
		node1, node2 = node1.Next, node2.Next
	}

	var rest *ListNode
	if node1 != nil {
		rest = node1
	}
	if node2 != nil {
		rest = node2
	}

	for rest != nil {
		val := rest.Val + carry
		if val >= 10 {
			val = val % 10
			carry = 1
		} else {
			carry = 0
		}
		current.Next = &ListNode{Val: val}
		current = current.Next
		rest = rest.Next
	}

	if carry == 1 {
		current.Next = &ListNode{Val: 1}
	}
	return root.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	current := root
	carry := 0

	node1, node2 := l1, l2
	for node1 != nil || node2 != nil {
		v1, v2 := 0, 0
		if node1 != nil {
			v1 = node1.Val
		}
		if node2 != nil {
			v2 = node2.Val
		}

		val := v1 + v2 + carry
		if val >= 10 {
			val = val % 10
			carry = 1
		} else {
			carry = 0
		}
		node := &ListNode{Val: val}

		current.Next = node
		current = current.Next

		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}

	if carry == 1 {
		current.Next = &ListNode{Val: 1}
	}
	return root.Next
}
