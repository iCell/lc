package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	cur := head
	var pre, next *ListNode

	for cur != nil {
		next = cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func reverseList1(head *ListNode) *ListNode {
	return reverseInner(nil, head)
}

func reverseInner(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverseInner(cur, next)
}
