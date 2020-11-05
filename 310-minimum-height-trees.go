func findMinHeightTrees(n int, edges [][]int) []int {
	nodes := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = make([]int, 0)
	}
	for _, edge := range edges {
		neighbors := nodes[edge[0]]
		neighbors = append(neighbors, edge[1])
		nodes[edge[0]] = neighbors
		
		another := nodes[edge[1]]
		another = append(another, edge[0])
		nodes[edge[1]] = another
	}

	min := math.MaxInt64
	results := make(map[int][]int)
	for i := 0; i < n; i++ {
		depth := helper(i, nodes)
		results[depth] = append(results[depth], i)
		if min > depth {
			min = depth
		}
	}

	return results[min]
}

func helper(node int, nodes map[int][]int) int {
	var depth int
	queue, visited := make([]int, 0), make(map[int]bool)
	queue = append(queue, node)
	visited[node] = true
	for len(queue) != 0 {
		size := len(queue)

		for size > 0 {
			first, temp := queue[0], queue[1:]
			queue = temp

			for _, neighbor := range nodes[first] {
				if visited[neighbor] == true {
					continue
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
			size--
		}
		depth++
	}
	return depth - 1
}
