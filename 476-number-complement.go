func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	
	c, n := 0, num
	for n > 0 {
		c += 1
		n = n >> 1
	}
	
	return num ^ int((math.Pow(2, float64(c)) - 1))
}