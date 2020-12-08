func numTrees(n int) int {
	dp := make([]int, n + 1, n + 1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i - 1 - j]
		}
	}
	return dp[n]
}

func numTrees(n int) int {
	return helper(1, n)
}

func helper(start, end int) int {
	if start >= end {
		return 1
	}
	
	var count int
	for i := start; i <= end; i++ {
		left := helper(start, i - 1)
		right := helper(i + 1, end)
		count += (left * right)
	}
	return count
}



