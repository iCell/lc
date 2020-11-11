func average(salary []int) float64 {
	min, max := math.MaxInt64, math.MinInt64
	var sum int
	for _, num := range salary {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
		sum += num
	}
	total := sum - min - max
	return float64(total) / float64(len(salary) - 2)
}