package main

func subsets(nums []int) [][]int {
	var result [][]int
	helper(&result, 0, len(nums), nums, nil)
	return result
}

func helper(result *[][]int, level, max int, nums, temp []int) {
	if level >= max {
		*result = append(*result, append([]int{}, temp...))
		return
	}

	helper(result, level+1, max, nums, temp)
	temp = append(temp, nums[level])
	helper(result, level+1, max, nums, temp)
	temp = temp[:len(temp)-1]
}