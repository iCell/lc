/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy, current := &ListNode{}, head
	
	for current != nil {
		pre := dummy
		
		for pre.Next != nil && pre.Next.Val < current.Val {
			pre = pre.Next
		}
		
		next := current.Next
		
		// insert to new list, between pre and pre.Next
		current.Next = pre.Next
		pre.Next = current
		
		current = next
	}
	
	return dummy.Next
}