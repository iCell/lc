package main

// stupid solution
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if last == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			last = nums[i]
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
