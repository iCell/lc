// f[n] = min(f[n-1], f[n-2]) + current
// f[0] = f[0]
// f[1] = f[1]

func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost), len(cost))
    dp[0] = cost[0]
    dp[1] = cost[1]

    for i := 2; i < len(cost); i++ {
        dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
    }

    return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

