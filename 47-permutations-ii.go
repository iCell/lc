package main

import "sort"

func main() {

}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return helper(nums)
}

func helper(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var result [][]int
	for i, num := range nums {
		if i != 0 && num == nums[i-1] {
			continue
		}

		rest := make([]int, len(nums)-1)
		copy(rest[:i], nums[:i])
		copy(rest[i:], nums[i+1:])

		for _, s := range helper(rest) {
			result = append(result, append(s, num))
		}
	}

	return result
}
