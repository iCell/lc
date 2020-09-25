// if s[i] == p[j] 或者 p[j] == '?', 则 dp[i][j] = dp[i-1][j-1]
// if p[j] == "*"，则氛围 * 匹配零个或者匹配多个的情况
//  // 1. * 匹配 0 个，则直接把当前的 * 消除，那么 dp[i][j] = dp[i][j-1]
//  // 1. * 匹配任意一个或者多个，则相当于消除 s[i]，然后看看前面的字符串匹配情况 dp[i][j] = dp[i-1][j]

// 若 s 或者 p 都为空字符串，则 dp[0][0] = true
// 若 s 为非空字符串，而 p 为空字符串，则 dp[i(i > 0)][0] = false
// 若 s 为空字符串，而 p 为非字符串，则当 j 全为 * 时，dp[0][j(j > 0)] = true

func isMatch(s string, p string) bool {
    sl, pl := len(s), len(p)

    dp := make([][]bool, sl+1, sl+1)
    for i := range dp {
        dp[i] = make([]bool, pl+1, pl+1)
    }

    dp[0][0] = true
    for j := 1; j <= pl; j++ {
        if p[j-1] != '*' {
            break
        }
        dp[0][j] = true
    }

    for i := 1; i <= sl; i++ {
        for j := 1; j <= pl; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '?' {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            }
        }
    }

    return dp[sl][pl]
}