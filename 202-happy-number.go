func isHappy(n int) bool {
	memo := make(map[int]bool)
	for n != 1 {
		if memo[n] == true {
			return false
		}
		memo[n] = true
		n = next(n)
	}
	return true
}

func next(n int) int {
	var result int
	for n > 0 {
		d := n % 10
		n = n / 10
		result += d * d
	}
	return result
}