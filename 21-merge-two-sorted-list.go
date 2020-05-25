package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	head1, head2 := l1, l2
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			next := head2.Next
			cur.Next = head2
			head2 = next
		} else {
			next := head1.Next
			cur.Next = head1
			head1 = next
		}
		cur = cur.Next
	}

	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}

	return dummy.Next
}
