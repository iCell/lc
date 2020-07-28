// 考察 s 中有多少个子序列同 t 相匹配，那么就往前推导：
// s[i-1] 中有多少个子序列同 t 匹配，f[i][j] = f[i-1][j]
// 这时再看 s[i] 和 t[i]，如果两者不想等，则没什么需要更改的
// 如果两者相等，则要把 s[i] 给算进去，因为前面已经考虑过 f[i-1][j] 的所有情况，又因为 s[i] 和 t[i] 相等，则剩余的情况就是 f[i-1][j-1]
func numDistinct(s string, t string) int {
    m, n := len(s)+1, len(t)+1

    dp := make([][]int, m, m)
    for i := range dp {
        dp[i] = make([]int, n, n)
        dp[i][0] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if s[i-1] == t[j-1] {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }

    return dp[m-1][n-1]
}