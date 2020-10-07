/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    slow, fast := head, head
    for k > 0 {
        if fast != nil && fast.Next != nil {
            fast = fast.Next
        } else {
            fast = head
        }
        k--
    }
    if slow == fast {
        return head
    }

    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }

    newHead := slow.Next
    fast.Next = head
    slow.Next = nil

    return newHead
}
