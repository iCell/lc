func canPermutePalindrome(s string) bool {
	letters := make(map[rune]int)
	for _, r := range s {
		letters[r] += 1
	}
	
	var count int
	for _, v := range letters {
		if v % 2 != 0 {
			count += 1
		}
		if count > 1 {
			return false
		}
	}
	return true
}