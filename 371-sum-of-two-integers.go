func getSum(a int, b int) int {
	for b != 0 {
		sum, carry := a^b, a&b << 1
		a, b = sum, carry
	}
	return a
}