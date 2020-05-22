package main

// stupid solution
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				if i != j {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}

// less stupid
func twoSum2(nums []int, target int) []int {
	temp := make(map[int]int, len(nums))
	for i, num := range nums {
		temp[num] = i
	}

	for i, num := range nums {
		another := target - num
		j, ok := temp[another]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	temp := make(map[int]int, len(nums))
	for i, num := range nums {
		j, ok := temp[target-num]
		if ok {
			return []int{i, j}
		}
		temp[num] = i
	}
	return nil
}
