func kLengthApart(nums []int, k int) bool {
	length := k
	for _, num := range nums {
		if num == 0 {
			length += 1
			continue
		}
		if length < k {
			return false
		}
		length = 0
	}
	return true
}