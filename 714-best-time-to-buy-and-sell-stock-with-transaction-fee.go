// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])

func maxProfit(prices []int, fee int) int {
    if len(prices) == 0 {
        return 0
    }

    dp := make([][]int, len(prices), len(prices))
    for i := range dp {
        dp[i] = make([]int, 2, 2)
    }
    dp[0][0], dp[0][1] = 0, -prices[0]

    var profit int
    for i := 1; i < len(dp); i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])

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

// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
func maxProfit2(prices []int, fee int) int {
    dp := make([][]int, len(prices), len(prices))
    for i := range dp {
        dp[i] = make([]int, 2, 2)
    }
    dp[0][0], dp[0][1] = 0, -prices[0]-fee

    var result int
    for i := 1; i < len(dp); i++ {
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        result = max(result, dp[i][0])
    }

    return result
}