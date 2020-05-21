package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	var (
		pre  *ListNode
		cur  = head
		next *ListNode
	)
	for {
		if cur == nil {
			return pre
		}
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseList2(head *ListNode) *ListNode {
	return reverseListRecursivelyInner(nil, head)
}

func reverseListRecursivelyInner(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverseListRecursivelyInner(cur, next)
}
