func isOneEditDistance(short string, long string) bool {
	if len(short) > len(long) {
		return isOneEditDistance(long, short)
	}
	if len(long)-len(short) > 1 {
		return false
	}
	for i := 0; i < len(short); i++ {
		if short[i] == long[i] {
			continue
		}
		if len(short) == len(long) {
			return short[i+1:] == long[i+1:]
		} else {
			return short[i:] == long[i+1:]
		}
	}
	return len(short)+1 == len(long)
}

