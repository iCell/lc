package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var (
		slow = head
		fast = head.Next.Next
	)

	for {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
}
