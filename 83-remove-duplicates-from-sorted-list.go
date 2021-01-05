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
		next := current.Next
		if pre != dummy && current.Val == pre.Val {
			pre.Next = next
		} else {
			pre = current
		}
		current = next
	}
	return dummy.Next
}