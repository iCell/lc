func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	v := float64(minutesToTest / minutesToDie + 1)
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(v)))
}