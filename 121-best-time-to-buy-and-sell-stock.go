package main

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			v := prices[j] - prices[i]
			if v > max {
				max = v
			}
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			p := prices[i] - minPrice
			if p > maxProfit {
				maxProfit = p
			}
		}
	}
	return maxProfit
}

// yes: 0, no: 1

// f[i][yes]: 代表第 i 天手上 有 股票
// f[i][no]: 代表第 i 天手上 没有 股票

// f[i][yes]:
// 1. i-1 就有股票，今天没有动
// 2. i-1 没有股票，但是 i 天有股票(这种情况在改题目下视为之前没有股票)
// f[i][yes] = max(f[i-1][yes], -prices[i])

// f[i][no]:
// 1. i-1 没有股票，今天也没有买
// 2. i-1 有股票，今天卖了
// f[i][no] = max(f[i-1][no], f[i-1][no] + prices[i])

// f[0][yes] = -prices[0], f[0][no] = 0

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices), len(prices))
	for i := range dp {
		dp[i] = make([]int, 2, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	result := dp[len(prices)-1][1]
	if result < 0 {
		result = 0
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
