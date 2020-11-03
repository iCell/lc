func maxPower(s string) int {
	if len(s) == 1 {
		return 1
	}
	
	max, cur, length := math.MinInt64, s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == cur {
			length += 1
		} else {
			cur, length = s[i], 1
		}
		if length > max {
			max = length
		}
	}
	return max
}