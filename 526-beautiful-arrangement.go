func countArrangement(n int) int {
	visited, count := make(map[int]bool), 0
	dfs(1, n, visited, &count)
	return count
}

func dfs(start, total int, visited map[int]bool, count *int) {
	if start > total {
		*count += 1
		return
	}
	for i := 1; i <= total; i++ {
		if visited[i] {
			continue
		}
		index, value := start, i
		if index % value == 0 || value % index == 0 {
			visited[i] = true
			dfs(start + 1, total, visited, count)
			visited[i] = false
		}
	}
}