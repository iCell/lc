/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	
	stack, current := make([]int, 0), head
	for current != nil {
		stack = append(stack, current.Val)
		current = current.Next
	}
	
	p1, p2 := len(stack) - 1, head
	for p1 >= 0 && p2 != nil {
		if stack[p1] != p2.Val {
			return false
		}
		p1--
		p2 = p2.Next
	}
	
	return true
}

func isPalindrome(head *ListNode) bool {
	var length int
	current := head
	for current != nil {
		length += 1
		current = current.Next
	}
	
	var pre *ListNode
	cur := head
	for i := 0; i < length / 2; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	if length % 2 != 0 {
		cur = cur.Next
	}
	for pre != nil && cur != nil {
		if pre.Val != cur.Val {
			return false
		}
		pre, cur = pre.Next, cur.Next
	}
	
	return true
}
