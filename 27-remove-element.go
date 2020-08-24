func removeElement(nums []int, val int) int {
	idx := 0
	for _, num := range nums {
		if num == val {
			continue
		}
		nums[idx] = num
		idx++
	}
	return idx
}