/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next: head}

    pre, node := dummy, head
    for node != nil {
        if node.Val == val {
            next := node.Next
            pre.Next = next

            node = next
            continue
        }
        pre = node
        node = node.Next
    }

    return dummy.Next
}