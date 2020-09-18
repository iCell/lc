func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		// 只能根据 mid 和 low 来判断是左边还是右边
		// 因为 low 是两段中的分界点，若选择 high，则 mid >= high 并不一定代表处于左边，因为 mid == high 就是在右边
		if nums[mid] >= nums[low] {
			if nums[mid] > target && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}