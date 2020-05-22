package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1 = reverseKGroup(node1, 2)

	n := node1
	for {
		if n == nil {
			return
		}
		fmt.Println(n.Val)
		n = n.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}

	return dummy.Next
}

func reverseBetweenNodes(left *ListNode, right *ListNode) *ListNode {
	var pre, next *ListNode
	cur := left.Next
	first := cur
	for {
		if cur == right {
			left.Next = pre
			first.Next = cur
			return pre
		}
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
}
