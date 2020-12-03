func shortestDistance(words []string, word1 string, word2 string) int {
	memo := make(map[string][]int)
	for i, word := range words {
		memo[word] = append(memo[word], i)
	}
	
	ans := math.MaxInt32
	
	idx1, idx2 := memo[word1], memo[word2]
	
	for i := range idx1 {
		for j := range idx2 {
			ans = min(ans, abs(idx1[i], idx2[j]))            
		}
	}
	
	return ans
}

func shortestDistance(words []string, word1 string, word2 string) int {
	i, j, ans := -1, -1, math.MaxInt32
	
	for idx, word := range words {
		if word == word1 {
			i = idx
		}
		if word == word2 {
			j = idx
		}
		if i != -1 && j != -1 {
			ans = min(ans, abs(i, j))
		}
	}
	
	return ans
}

