// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])

// dp[0][0], dp[0][1] = 0, -prices[0]
// dp[1][0], dp[1][1] = max(0, dp[0][1] + prices[1]), max(-prices[0], -prices[1])

func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    dp := make([][]int, len(prices), len(prices))
    for i := range dp {
        dp[i] = make([]int, 2, 2)
    }

    dp[0][0], dp[0][1] = 0, -prices[0]
    dp[1][0], dp[1][1] = max(0, dp[0][1]+prices[1]), max(-prices[0], -prices[1])

    profit := max(dp[0][0], dp[1][0])
    for i := 2; i < len(prices); i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])

        profit = max(profit, dp[i][0])
    }

    return profit
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}