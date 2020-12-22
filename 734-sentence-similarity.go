func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	
	s := make(map[string]bool)
	for _, pair := range similarPairs {
		s[pair[0] + "#" + pair[1]] = true
	}
	
	for index := 0; index < len(sentence1); index++ {
		left, right := sentence1[index], sentence2[index]
		if left == right {
			continue
		}
		if s[left + "#" + right] == false && s[right + "#" + left] == false {
			return false
		}
	}
	return true
}