// dp[i][j] 代表 s[i](s 第 i-1 个字符) 和 p[j]（p 第 j-1 个字符） 以及之前的字符下能不能匹配成功
// 若 s[i] == p[j] 或者 p[j] == '.' 则能够匹配成功
// 若 p[j] 为 '*'，因此 * 要配合前面的字符使用:
// // 1. 若 s[i] 和 p[j-1] 匹配，则代表要不是两个字符相同，要不就是 p[j-1] 为 '.'
// // // 1.1 若匹配 0 次，则直接将 p[j-1] + * 两个字符丢弃即可，dp[i][j] = dp[i][j-2]
// // // 1.2 若匹配 1 次，则相当于把 s[i] 以及 p[j-1] + * 给抛弃掉，那么 dp[i][j] = dp[i-1][j-2]
// // // 1.3 若匹配大于 1 次，则先把 s[i] 给匹配掉抹去，再接着看，也就是 dp[i][j] = dp[i-1][j]
// // 2. 若 s[i] 和 p[j-1] 不匹配，那就相当于 1.1 的情况，把 p[j-1] + * 舍弃掉

// 若 s 和 p 都为空字符串，dp[0][0] = true
// 若 s 非空，而 p 为空，则肯定匹配不上，则 dp[i(i > 0)][0] = false
// 若 s 空，而 p 为非空，则只可能类似于 c* 这样的模式，让 c 能够被 * 干掉，所以当 p[j-1] 为 * 且 dp[0][j] = dp[0][j-2]

func isMatch(s string, p string) bool {
    sl, pl := len(s), len(p)
    dp := make([][]bool, sl+1, sl+1)
    for i := range dp {
        dp[i] = make([]bool, pl+1, pl+1)
    }

    dp[0][0] = true
    for j := 1; j <= pl; j++ {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-2]
        }
    }

    for i := 1; i <= sl; i++ {
        for j := 1; j <= pl; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                if s[i-1] == p[j-2] || p[j-2] == '.' {
                    dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-2]
                }
            }
        }
    }

    return dp[sl][pl]
}