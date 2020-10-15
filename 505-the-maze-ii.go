var directions = [][]int{
    []int{0, 1},
    []int{0, -1},
    []int{1, 0},
    []int{-1, 0},
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
    m, n := len(maze), len(maze[0])

    queue := make([][]int, 0)
    queue = append(queue, start)

    distance := make([][]int, m, m)
    for i := range distance {
        distance[i] = make([]int, n, n)
        for j := range distance[i] {
            distance[i][j] = -1
        }
    }
    distance[start[0]][start[1]] = 0

    for len(queue) != 0 {
        first, temp := queue[0], queue[1:]
        queue = temp

        for _, direction := range directions {
            steps := distance[first[0]][first[1]]

            newX, newY := first[0]+direction[0], first[1]+direction[1]
            for newX >= 0 && newX < m && newY >= 0 && newY < n && maze[newX][newY] == 0 {
                newX += direction[0]
                newY += direction[1]
                steps++
            }

            destX, destY := newX-direction[0], newY-direction[1]
            if distance[destX][destY] == -1 || distance[destX][destY] > steps {
                queue = append(queue, []int{destX, destY})
                distance[destX][destY] = steps
            }
        }
    }

    return distance[destination[0]][destination[1]]
}