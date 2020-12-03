func findMaxConsecutiveOnes(nums []int) int {
	var result int
	
	var i int
	for i < len(nums) {
		if nums[i] == 0 {
			i++
			continue
		}
		j := i
		for j < len(nums) && nums[j] == 1 {
			j++
		}
		if j - i > result {
			result = j - i
		}
		i = j
	}
	
	return result
}