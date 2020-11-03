/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		pre, next := dummy, head.Next    
		for pre.Next != nil && head.Val > pre.Next.Val {
			pre = pre.Next
		}
		head.Next = pre.Next
		pre.Next = head
		head = next
	}
	return dummy.Next
}
