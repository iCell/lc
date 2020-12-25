package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}

	index := 0
	node, left := dummy, dummy
	for node != nil {
		if index != 0 && index%2 == 0 {
			left = swapBetween(left, node.Next)
			node = left.Next
		} else {
			node = node.Next
		}

		index++
	}

	return dummy.Next
}

func swapBetween(left, right *ListNode) *ListNode {
	pre, cur := left, left.Next
	first := cur
	var next *ListNode
	for cur != right {
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	left.Next = pre
	first.Next = right

	return first
}
