func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	a := strings.Join(word1, "")
	b := strings.Join(word2, "")
	return a == b
}