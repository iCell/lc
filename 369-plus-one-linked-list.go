/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
	stack := make([]*ListNode, 0)
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}
	
	dummy := &ListNode{}
	current, carry := dummy, 1
	
	for len(stack) != 0 {
		pop, temp := stack[len(stack)-1], stack[:len(stack)-1]
		stack = temp
		
		val := pop.Val + carry
		if val > 9 {
			val = val % 10
			carry = 1
		} else {
			carry = 0
		}
		
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{Val: 1}
		current = current.Next
	}
	
	result := reverse(nil, dummy.Next)
	
	return result
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	return reverse(cur, next)
}