type WordDistance struct {
	Memo map[string][]int    
}


func Constructor(words []string) WordDistance {
	memo := make(map[string][]int)
	for i, word := range words {
		memo[word] = append(memo[word], i)
	}
	return WordDistance{
		Memo: memo,
	}
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
	ans := math.MaxInt32
	
	idx1, idx2 := this.Memo[word1], this.Memo[word2]

	i, j := 0, 0
	for i < len(idx1) && j < len(idx2) {
		ans = min(ans, abs(idx1[i], idx2[j]))
		if idx1[i] < idx2[j] {
			i++
		} else {
			j++
		}
	}
	
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */