func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x % 10 == 0) {
		return false
	}
	if x < 10 {
		return true
	}

	var reverted int
	for x > reverted {
		reverted = reverted * 10 + x % 10
		x = x / 10
	}

	return x == reverted || x == reverted / 10
}