package main

import "math"

func canJump(nums []int) bool {
	max := 0
	for i, num := range nums {
		if i > max {
			return false
		}
		max = int(math.Max(float64(max), float64(i+num)))
	}
	return true
}
