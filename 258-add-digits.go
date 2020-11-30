func addDigits(num int) int {
	for num >= 10 {
		num = next(num)
	}
	return num
}

func next(num int) int {
	var ans int
	for num > 0 {
		ans += num % 10
		num = num / 10
	}
	return ans
}