func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodes := make(map[int]*ListNode)

    var index int

    node := head
    for node != nil {
        nodes[index] = node
        node = node.Next
        index++
    }

    idx := index - n
    if idx == 0 {
        return head.Next
    }

    last, current := nodes[idx-1], nodes[idx]
    last.Next = current.Next

    return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
    slow, fast := head, head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    if fast == nil {
        return slow.Next
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next

    return head
}
