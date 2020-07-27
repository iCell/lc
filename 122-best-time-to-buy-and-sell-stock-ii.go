package main

func maxProfit(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			max += diff
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	dp := make([][]int, len(prices), len(prices))
	for i := range dp {
		dp[i] = make([]int, 2, 2)
	}
	dp[0][0], dp[0][1] = 0, -prices[0]

	for i := 1; i < len(dp); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
