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
