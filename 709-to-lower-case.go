func toLowerCase(str string) string {
	// 'a' - 'A' = 32
	runes := []rune(str)
	for i, r := range runes {
		if r >= 65 && r <= 90 {
			runes[i] = r + 32
		}
	}
	return string(runes)
}