package main

import "fmt"

func main() {
	fmt.Println(permute2([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var result [][]int
	for i, num := range nums {
		rest := make([]int, len(nums)-1)
		copy(rest[:i], nums[:i])
		copy(rest[i:], nums[i+1:])

		for _, r := range permute(rest) {
			result = append(result, append(r, num))
		}
	}
	return result
}

func permute2(nums []int) [][]int {
	result := make([][]int, 0)
	var temp []int
	helper(&result, nums, temp)
	return result
}

func helper(result *[][]int, nums []int, temp []int) {
	if len(nums) == len(temp) {
		*result = append(*result, append([]int{}, temp...))
		return
	}

	for _, num := range nums {
		contains := false
		for _, t := range temp {
			if t == num {
				contains = true
				break
			}
		}
		if contains {
			continue
		}

		temp = append(temp, num)
		helper(result, nums, temp)
		temp = temp[:len(temp)-1]
	}
}

func permute(nums []int) [][]int {
	temp, result := make([]int, 0), make([][]int, 0)
	visited := make(map[int]bool)
	helper(nums, temp, &result, visited)
	return result
}

func helper(nums []int, permutation []int, result *[][]int, visited map[int]bool) {
	if len(nums) == len(permutation) {
		*result = append(*result, append([]int{}, permutation...))
		return
	}

	for _, num := range nums {
		if visited[num] == true {
			continue
		}

		permutation = append(permutation, num)
		visited[num] = true

		helper(nums, permutation, result, visited)

		permutation = permutation[:len(permutation)-1]
		delete(visited, num)
	}
}
