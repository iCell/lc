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

func helper2(nums []int, index int, current []int, result *[][]int) {
	*result = append(*result, append([]int{}, current...))

	for i := index; i < len(nums); i++ {
		// [1] -> [1, 2]
		current = append(current, nums[i])
		// [1, 2] -> [1, 2, 3]
		// [1, 2] -> [1, 2, 4]
		helper(nums, i+1, current, result)
		// [1, 2] -> [1]s
		current = current[:len(current)-1]
	}
}
