// 第 i 天，还剩下 k 次可以买，手里有股票：dp[i][k][1]
// i-1 天，手里就有股票 dp[i-1][k][1]
// i-1 天，手里没有股票，第 i 天买了股票 dp[i-1][k-1][0] - price[i]
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

// 第 i 天，已经买了 k 次，手里没有股票：dp[i][k][0]
// i-1 天，手里就没有股票 dp[i-1][k][0]
// i-1 天，手里有股票，第 i 天卖了股票 dp[i-1][k][1] + price[i]
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])

// 最大购买次数为 2 次，则 1 <= k <= 2, k = 0 代表尚未开始交易
// dp[0][0][0] = 0
// dp[0][0][1] 尚未开始交易，手中不可能持有股票

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    maxK := 2
    dp := make([][][]int, len(prices), len(prices))
    for i := range dp {
        dp[i] = make([][]int, maxK+1, maxK+1)
        for k := range dp[i] {
            dp[i][k] = make([]int, 2, 2)
        }
    }

    var profit int
    for i := 0; i < len(dp); i++ {
        for k := maxK; k > 0; k-- {
            if i == 0 {
                dp[0][k][1] = -prices[0]
            } else {
                dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
                dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
            }
            profit = max(profit, dp[i][k][0])
        }
    }

    return profit
}

// dp[i][k] = Max(dp[i-1][k], prices[i] - prices[j] + dp[j][k-1])，j 取 0 - i
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    dp := make([][]int, len(prices), len(prices))
    for i := range dp {
        dp[i] = make([]int, 3, 3)
    }

    for i := 1; i < len(prices); i++ {
        for k := 2; k > 0; k-- {
            var temp int
            for j := 0; j <= i; j++ {
                temp = max(temp, prices[i]-prices[j]+dp[j][k-1])
            }
            dp[i][k] = max(temp, dp[i-1][k])
        }
    }

    var profit int
    for k := 2; k > 0; k-- {
        profit = max(profit, dp[len(dp)-1][k])
    }

    return profit
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}