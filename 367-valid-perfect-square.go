package main

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left, right := 0, num
	for left <= right {
		mid := left + (right-left)>>1
		result := mid * mid
		if result == num {
			return true
		} else if result > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
