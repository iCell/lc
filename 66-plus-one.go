package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}

func plusOne(digits []int) []int {
	plus := true
	length := len(digits) + 1
	result := make([]int, length, length)
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i]
		if plus {
			if v == 9 {
				result[i+1] = 0
				plus = true
			} else {
				result[i+1] = v + 1
				plus = false
			}
		} else {
			result[i+1] = v
		}
	}
	if plus {
		result[0] = 1
		return result
	}
	return result[1:]
}
