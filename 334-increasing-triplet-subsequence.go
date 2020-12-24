func increasingTriplet(nums []int) bool {
	small, big := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if small >= num {
			small = num
		} else if big >= num {
			big = num
		} else {
			return true
		}
	}
	return false
}