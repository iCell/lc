func longestValidParentheses(s string) int {
    maxLength := 0

    stack := []int{-1}
    for i, r := range s {
        if r == rune('(') {
            stack = append(stack, i)
            continue
        }
        stack = stack[:len(stack)-1]
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            length := i - stack[len(stack)-1]
            if length > maxLength {
                maxLength = length
            }
        }
    }

    return maxLength
}
