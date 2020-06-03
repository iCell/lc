package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{nums}
	}

	var result [][]int
	for i, num := range nums {
		rest := make([]int, len(nums)-1)
		copy(rest[:i], nums[:i])
		copy(rest[i:], nums[i+1:])

		for _, s := range permute(rest) {
			result = append(result, append(s, num))
		}
	}

	return result
}
