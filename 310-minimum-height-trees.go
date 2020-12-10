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

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	
	graph, degree := make(map[int][]int, n), make([]int, n)
	for i := range edges {
		u, v := edges[i][0], edges[i][1]
		
		graph[v] = append(graph[v], u)
		graph[u] = append(graph[u], v)
		
		degree[v] += 1
		degree[u] += 1
	}
	
	queue := make([]int, 0)
	for i := 0; i < len(degree); i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	
	for len(queue) < n {
		size := len(queue)
		n -= size
		for size > 0 {
			first, temp := queue[0], queue[1:]
			queue = temp
			
			for _, node := range graph[first] {
				degree[node] -= 1
				if degree[node] == 1 {
					queue = append(queue, node)
				}
			}
			
			size -= 1
		}
	}
	
	return queue
}


