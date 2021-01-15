func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	arr, ans := make([]int, n + 1), math.MinInt32
	arr[0], arr[1] = 0, 1
	for i := 2; i <= n; i++ {
		d := (i - i % 2) / 2
		if i % 2 == 0 {
			arr[i] = arr[d]
		} else {
			arr[i] = arr[d] + arr[d+1]
		}
		ans = max(ans, arr[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}