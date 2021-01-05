/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	memo, dummy := make(map[int]bool), &ListNode{Next: head}
	pre, current := dummy, head
	for current != nil {
		next := current.Next
		if memo[current.Val] {
			pre.Next = next
		} else {
			pre = current
		}
		memo[current.Val] = true
		current = next
	}
	return dummy.Next
}