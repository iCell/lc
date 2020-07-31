func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    if len(matrix[0]) == 0 {
        return 0
    }

    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m, m)
    for i := range dp {
        dp[i] = make([]int, n, n)
        for j := range dp[i] {
            if matrix[i][j] == '1' {
                dp[i][j] = 1
            }
        }
    }

    var result int
    for i := range dp[0] {
        result = max(result, dp[0][i])
    }
    for i := range dp {
        result = max(result, dp[i][0])
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if dp[i][j] == 1 {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
            result = max(dp[i][j], result)
        }
    }

    return result * result
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y, z int) int {
    m := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    return m(x, m(y, z))
}