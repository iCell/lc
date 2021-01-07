func countComponents(n int, edges [][]int) int {
	nodes := make(map[int][]int)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}
	
	var ans int
	
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		
		queue := make([]int, 0)
		queue = append(queue, i)
		
		for len(queue) != 0 {
			first, temp := queue[0], queue[1:]
			queue = temp
			
			for _, neighbor := range nodes[first] {
				if visited[neighbor] {
					continue
				}
				queue, visited[neighbor] = append(queue, neighbor), true
			}
		}
		
		ans += 1
	}
	
	return ans
}