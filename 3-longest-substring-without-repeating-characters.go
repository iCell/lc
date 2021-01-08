func lengthOfLongestSubstring(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		memo, length := make(map[byte]bool), 0
		for j := i; j < len(s); j++ {
			if memo[s[j]] {
				break
			}
			memo[s[j]], length = true, length + 1
		}
		ans = max(ans, length)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}