func countCharacters(words []string, chars string) int {
	memo := make(map[rune]int, len(chars))
	for _, c := range chars {
		memo[c] += 1
	}
	
	var ans int
	for _, word := range words {
		temp := make(map[rune]int, len(word))
		for _, c := range word {
			temp[c] += 1
		}
		
		good := true
		for k, v := range temp {
			if v > memo[k] {
				good = false
				break
			}
		}
		if good {
			ans += len(word)
		}
	}
	return ans
}