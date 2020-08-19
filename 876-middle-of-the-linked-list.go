func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for {
        if fast == nil {
            break
        }
        if fast.Next == nil {
            break
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}