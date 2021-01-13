// range [2^31 + 1, 2^31 - 1]
// if -2^31 / -1 = 2^31, which is overflow

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}    
	
	negative := 2
	if dividend > 0 {
		dividend = -dividend
		negative -= 1
	}
	if divisor > 0 {
		divisor = -divisor
		negative -= 1
	}
	
	var quotient int
	for dividend - divisor <= 0 {
		quotient--
		dividend -= divisor
	}
	
	if negative != 1 {
		quotient = -quotient
	}
	return quotient
}