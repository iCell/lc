func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, p := range position {
		if p % 2 == 0 {
			even++
			continue
		}
		odd++
	}
	return min(odd, even)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}