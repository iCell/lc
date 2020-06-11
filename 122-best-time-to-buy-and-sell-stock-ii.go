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
