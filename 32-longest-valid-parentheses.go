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

// 1. s[i]是 '('，以它为结尾的子串肯定不是有效括号子串: dp[i] = 0
// 2. s[i]是')'，前一个子问题的末尾 s[i-1]

// 2.1 s[i-1]是'('，它们俩结成一对( )，考察s[i-2]
// 2.1.1 s[i-2]不存在 dp[i] = 2
// 2.1.2 存在，dp[i] = dp[i-2] + 2

// 2.2 s[i-1]是')'，即 '))' 形式，这就来到子问题了，以 s[i-1] 为结尾形成的最长有效长度为 dp[i-1]，跨过这个长度，来看 s[i-dp[i-1]-1] 这个字符

// 如何理解：
// dp[i-1] 是 i 位前一位的最长有效括号的长度，那么 i - dp[i-1] 就是 i - 1 位内部最长有效括号的最左边的位置，然后 i - dp[i-1] - 1 就是跟 i 位上的 ) 相匹配的 ( 位置

// 2.2.1 s[i-dp[i-1]-1] 不存在或为 ')'，s[i]找不到匹配，dp[i] = 0
// 2.2.2 s[i-dp[i-1]-1] 是'('，它和s[i]呼应，有效长度 2 保底，加上跨过的 dp[i-1]，再加上前方的 dp[i-dp[i-1]-2]
// 2.2.2.1 s[i-dp[i-1]-2] 存在，dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
// 2.2.2.2 s[i-dp[i-1]-2]不存在，dp[i] = dp[i-1] + 2

// https://www.youtube.com/watch?v=39CEPFCl5sE
