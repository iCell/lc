package main

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for i := 2; i < 6; i++ {
		for num%i == 0 {
			num = num / i
		}
	}
	return num == 1
}

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
	return true
}
