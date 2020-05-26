package main

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	temp := make(map[int]int, len(nums))
	for i, num := range nums {
		j, exist := temp[target-num]
		if exist {
			return []int{i, j}
		}
		temp[num] = i
	}
	return nil
}
