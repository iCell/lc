var directions = [][]int {
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
	[]int{-1, 0},
}

func generateMatrix(n int) [][]int {
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	direction, x, y := 0, 0, 0

	for i := 0; i < n * n; i++ {
		visited[x][y] = i + 1

		newx := x + directions[direction][0]
		newy := y + directions[direction][1]

		outx := newx >= n || newx < 0
		outy := newy >= n || newy < 0
		if outx || outy || visited[newx][newy] != -1 {
			direction = (direction + 1) % len(directions)
		}

		x += directions[direction][0]
		y += directions[direction][1]
	}

	return visited
}
