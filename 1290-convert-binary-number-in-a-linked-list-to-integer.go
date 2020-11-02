/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	current := reverse(nil, head)
	var index int
	var result int
	
	for current != nil {
		result += current.Val * int(math.Pow(2.0, float64(index)))
		
		current = current.Next
		index += 1
	}
	
	return result
}

func reverse(pre, current *ListNode) *ListNode {
	if current == nil {
		return pre
	}
	next := current.Next
	current.Next = pre
	return reverse(current, next)
}