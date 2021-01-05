/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre, current := dummy, head
	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}   
			next := current.Next
			pre.Next = next
			current = next
		} else {
			next := current.Next
			pre = current
			current = next
		}
	}
	return dummy.Next
}