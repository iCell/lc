
// f[0][0] = current
// f[i][0] = f[i-1][0] + current
// f[i][i] = f[i-1][j-1] + current
// f[i][j] = min(f[i-1][j-1], f[i-1][j]) + current
func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }

    dp := make([][]int, len(triangle), len(triangle))
    for i := range dp {
        dp[i] = make([]int, len(triangle[i]), len(triangle[i]))
    }
    dp[0][0] = triangle[0][0]

    for i := 1; i < len(triangle); i++ {
        for j := 0; j < len(triangle[i]); j++ {
            current := triangle[i][j]
            if j == 0 {
                dp[i][j] = dp[i-1][0] + current
            } else if j == len(triangle[i])-1 {
                dp[i][j] = dp[i-1][j-1] + current
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + current
            }
        }
    }

    result := dp[len(triangle)-1][0]
    for i := 1; i < len(dp[len(triangle)-1]); i++ {
        result = min(result, dp[len(triangle)-1][i])
    }
    return result
}

// f[n][i] = min(f[n+1][i], f[n+1][i+1]) + current
// f[max][i] = f[max][i]
func minimumTotal2(triangle [][]int) int {
    if len(triangle) == 0 {
        return 0
    }

    dp := make([][]int, len(triangle), len(triangle))
    for i := range dp {
        dp[i] = make([]int, len(triangle[i]), len(triangle[i]))
    }

    for i := range dp[len(dp)-1] {
        dp[len(dp)-1][i] = triangle[len(dp)-1][i]
    }

    for i := len(dp) - 2; i >= 0; i-- {
        for j := 0; j <= len(dp[i])-1; j++ {
            dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
        }
    }

    return dp[0][0]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}


