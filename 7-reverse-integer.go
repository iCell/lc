func reverse(x int) int {
	var result int
	
	max, min := math.MaxInt32, math.MinInt32
	
	for x != 0 {
		if result > (max / 10) || result < (min / 10) {
			return 0
		}
		result = result * 10 + x % 10
		x = x/10
	}
	
	return result
}