32 位数字取值范围为 [-2^31, 2^31-1]
1. 若是 -2^31 / -1，则结果溢出了，因此要特殊处理
2. 若是  dividend 或者 dividor 中有一方是负数，则需要取负数，但此时存在一种情况，对 -2^31 取负数在某些语言中（比如 C，golang 不会）会越界

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	
	positive := 2
	if dividend < 0 {
		dividend = -dividend
		positive -= 1
	}
	if divisor < 0 {
		divisor = -divisor
		positive -= 1
	}
	
	var quotient int
	for divisor <= dividend {
		power, value := 1, divisor
		for value <= math.MaxInt32 / 2 && value + value <= dividend {
			value += value
			power += power
		}
		quotient += power
		dividend -= value
	}
	
	if positive == 1 {
		return -quotient
	}
	return quotient
}