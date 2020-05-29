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
	if len(nums) <= 1 {
		return len(nums)
	}
	var slowIndex, fastIndex int
	for fastIndex = 0; fastIndex < len(nums); fastIndex++ {
		slow, fast := nums[slowIndex], nums[fastIndex]
		if slow != fast {
			slowIndex += 1
			nums[slowIndex] = nums[fastIndex]
		}
	}
	return slowIndex + 1
}
