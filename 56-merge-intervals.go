func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i, result := 0, make([][]int, 0)
	for i < len(intervals) {
		best := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= best {
			best = max(best, intervals[j][1])
			j += 1
		}
		result = append(result, []int{intervals[i][0], best})
		i = j
	}
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}