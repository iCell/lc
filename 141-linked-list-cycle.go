package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	cur := head
	jum := head

	for {
		cur = cur.Next
		if jum.Next == nil {
			return false
		}
		jum = jum.Next.Next

		if jum == nil {
			return false
		}
		if cur.Val == jum.Val {
			return true
		}
	}
}
