func numJewelsInStones(J string, S string) int {
	hash := make(map[rune]int, len(S))
	for _, s := range S {
		hash[s] += 1
	}
	var count int
	for _, j := range J {
		count += hash[j]
	}
	return count
}