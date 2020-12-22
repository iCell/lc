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

