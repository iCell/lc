/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	
	mid := (left + right) / 2
	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)
	return mergeTwo(l1, l2)
}

func mergeTwo(first, second *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	
	node1, node2 := first, second
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			current.Next = node1
			node1 = node1.Next
		} else {
			current.Next = node2
			node2 = node2.Next
		}
		current = current.Next
	}
	if node1 != nil {
		current.Next = node1
	}
	if node2 != nil {
		current.Next = node2
	}
	return dummy.Next
}