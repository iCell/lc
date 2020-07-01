package main

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	return helper([]rune(text1), []rune(text2), len(text1)-1, len(text2)-1)
}

func helper(r1, r2 []rune, li, ri int) int {
	if li == -1 || ri == -1 {
		return 0
	}
	if r1[li] == r2[ri] {
		return helper(r1, r2, li-1, ri-1) + 1
	}
	return int(math.Max(float64(helper(r1, r2, li, ri-1)), float64(helper(r1, r2, li-1, ri))))
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1, m+1)
	for i := range dp {
		arr := make([]int, n+1, n+1)
		dp[i] = arr
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
