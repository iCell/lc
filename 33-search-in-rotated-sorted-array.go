package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if target < nums[0] {
			if nums[mid] < target || nums[mid] >= nums[0] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] < target && nums[mid] >= nums[0] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
