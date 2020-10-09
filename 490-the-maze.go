var directions = [][]int{
    []int{0, 1},
    []int{0, -1},
    []int{1, 0},
    []int{-1, 0},
}

func hasPath(maze [][]int, start []int, destination []int) bool {
    m, n := len(maze), len(maze[0])

    queue, visited := make([][]int, 0), make([][]bool, m, m)
    for i := range visited {
        visited[i] = make([]bool, n, n)
        if i == start[0] {
            visited[start[0]][start[1]] = true
        }
    }
    queue = append(queue, start)

    for len(queue) != 0 {
        first, temp := queue[0], queue[1:]
        queue = temp

        if first[0] == destination[0] && first[1] == destination[1] {
            return true
        }

        for _, direction := range directions {
            newX, newY := first[0]+direction[0], first[1]+direction[1]
            for newX >= 0 && newX < m && newY >= 0 && newY < n && maze[newX][newY] == 0 {
                newX += direction[0]
                newY += direction[1]
            }

            destX, destY := newX-direction[0], newY-direction[1]
            if visited[destX][destY] == true {
                continue
            }

            visited[destX][destY] = true
            queue = append(queue, []int{destX, destY})
        }
    }

    return false
}