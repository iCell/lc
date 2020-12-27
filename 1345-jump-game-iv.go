func minJumps(arr []int) int {
	memo := make(map[int][]int)
	for i, num := range arr {
		memo[num] = append(memo[num], i)
	}
	
	queue, visited := make([]int, 0), make(map[int]bool)
	queue, visited[0] = append(queue, 0), true
	
	var depth int
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			first, temp := queue[0], queue[1:]
			queue = temp
			
			if first == len(arr) - 1 {
				return depth
			}
			for _, v := range memo[arr[first]] {
				if visited[v] == true {
					continue
				}
				queue, visited[v] = append(queue, v), true
			}
			
			// this line is important
			delete(memo, arr[first])
			
			if first > 0 && visited[first-1] == false {
				queue, visited[first-1] = append(queue, first-1), true
			}
			if first < len(arr) - 1 && visited[first+1] == false {
				queue, visited[first+1] = append(queue, first+1), true
			}
			size -= 1
		}
		depth += 1
	}
	
	return depth
}