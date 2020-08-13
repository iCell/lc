// if s[start] == s[end]:
// dp[start][end] = dp[start+1][end-1] + 2
// else
// dp[start][end] = max(dp[start+1][end], dp[start][end-1])

func longestPalindromeSubseq(s string) int {
    if len(s) == 0 {
        return 0
    }

    dp := make([][]int, len(s), len(s))
    for i := range dp {
        dp[i] = make([]int, len(s), len(s))
        dp[i][i] = 1
    }

    for i := len(s) - 1; i >= 0; i-- {
        for j := i + 1; j < len(s); j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
                continue
            }
            dp[i][j] = max(dp[i+1][j], dp[i][j-1])
        }
    }

    return dp[0][len(s)-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
