// 重复则在哈希表中删去start，并将start后移；由于此时end未变所以下一次循环时判断的依旧是这个刚刚被判定为已经存在的数，直到与他相同的数被删去为止

func lengthOfLongestSubstring(s string) int {
	ans, runes := 0, []rune(s)
	i, j, memo := 0, 0, make(map[rune]bool)
	for i < len(runes) && j < len(runes) {
		_, exist := memo[runes[j]]
		if !exist {
			memo[runes[j]] = true
			j += 1
			ans = max(ans, j - i)
		} else {
			delete(memo, runes[i])
			i += 1
		}
	}
	return ans
}

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