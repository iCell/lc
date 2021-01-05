func pivotIndex(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	
	var left int
	for i, n := range nums {
		right := sum - left - n
		if left == right {
			return i
		}
		left += n
	}
	
	return -1
}