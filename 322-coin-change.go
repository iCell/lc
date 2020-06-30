package main

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]float64, amount+1, amount+1)
	for idx := range dp {
		dp[idx] = math.MaxFloat64
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := 0; i < amount+1; i++ {
			if i-coin < 0 {
				continue
			}
			dp[i] = math.Min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == math.MaxFloat64 {
		return -1
	}
	return int(dp[amount])
}
