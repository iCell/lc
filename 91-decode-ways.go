package main

func numDecodings(s string) int {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}

	dp := make([]int, len(runes), len(runes))
	if runes[0] != '0' {
		dp[0] = 1
	}

	for i := 1; i < len(runes); i++ {
		first := int(runes[i] - '0')
		if first > 0 {
			dp[i] += dp[i-1]
		}
		second := int(runes[i-1]-'0')*10 + int(runes[i]-'0')
		if second >= 10 && second <= 26 {
			if i >= 2 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}

	return dp[len(runes)-1]
}
