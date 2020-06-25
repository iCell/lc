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
