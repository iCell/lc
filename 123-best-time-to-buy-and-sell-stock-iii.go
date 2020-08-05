import "math"

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

// k 代表还剩下几次购买机会，那么总共能够进行买卖的次数为 k*2
// dp[i][k]
// if k != 0 && k % 2 == 1: max(dp[i-1][k], dp[i-1][k-1] - prices[i])
// if k != 0 && k % 2 == 0: max(dp[i-1][k], dp[i-1][k-1] + prices[i])
//
// dp[0][0] = 0, dp[0][1] = -prices[0]

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    times := 2

    dp := make([][]int, len(prices), len(prices))
    for i := range dp {
        dp[i] = make([]int, times*2+1, times*2+1)
    }
    for k := 1; k <= times*2; k++ {
        if k%2 == 1 {
            dp[0][k] = -prices[0]
        }
    }

    for i := 1; i < len(dp); i++ {
        for k := 0; k <= times*2; k++ {
            if k%2 == 0 {
                if k == 0 {
                    dp[i][0] = dp[i-1][0]
                } else {
                    dp[i][k] = max(dp[i-1][k], dp[i-1][k-1]+prices[i])
                }
            } else {
                dp[i][k] = max(dp[i-1][k], dp[i-1][k-1]-prices[i])
            }
        }
    }

    var profit int
    for i := range dp[len(dp)-1] {
        if i%2 == 0 {
            profit = max(profit, dp[len(dp)-1][i])
        }
    }

    return profit
}

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    times := 2*2 + 1

    dp := make([]int, times, times)
    for i := range dp {
        dp[i] = math.MinInt64
    }
    dp[0], dp[1] = 0, -prices[0]

    var profit int
    for i := 1; i < len(prices); i++ {
        for k := 1; k < times; k++ {
            switch {
            case k%2 == 1:
                dp[k] = max(dp[k], dp[k-1]-prices[i])
            case k%2 == 0:
                dp[k] = max(dp[k], dp[k-1]+prices[i])
                profit = max(dp[k], profit)
            }
        }
    }

    return profit
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}