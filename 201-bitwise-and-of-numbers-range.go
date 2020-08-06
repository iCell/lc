// m 到 n 是连续的，所有它们的二进制表示中，从最低位开始一定有一个数是 0，而最终的结果就是它们的最高位肯定是 1
func rangeBitwiseAnd(m int, n int) int {
	var zeros int
	for m < n {
		zeros++
		m = m >> 1
		n = n >> 1
	}
	return m << zeros
}