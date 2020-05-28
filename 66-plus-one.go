package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	plus := true
	for i := len(digits) - 1; i >= 0; i-- {
		if !plus {
			continue
		}
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
			plus = true
		} else {
			plus = false
		}
	}
	if plus {
		return append([]int{1}, digits...)
	}
	return digits
}
