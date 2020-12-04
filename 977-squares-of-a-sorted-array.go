func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	
	ans := make([]int, len(nums))
	
	left, right := 0, len(nums) - 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		if nums[left] * nums[left] > nums[right] * nums[right] {
			ans[idx] = nums[left] * nums[left]
			left += 1
		} else {
			ans[idx] = nums[right] * nums[right]
			right -= 1
		}
	}
	
	return ans
}

