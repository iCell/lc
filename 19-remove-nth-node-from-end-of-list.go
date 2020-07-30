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