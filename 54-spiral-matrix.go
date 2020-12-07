var directions = [][]int {
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
	[]int{-1, 0},
}

func spiralOrder(matrix [][]int) []int { 
	rows, columns := len(matrix), len(matrix[0])
	if rows == 0 || columns == 0 {
		return nil
	}

	visited := make([][]bool, rows, rows)
	for i := range visited {
		visited[i] = make([]bool, columns, columns)
	}

	direction, x, y := 0, 0, 0

	var ans []int

	for i := 0; i < rows * columns; i++ {
		visited[x][y] = true
		ans = append(ans, matrix[x][y])

		newx := x + directions[direction][0]
		newy := y + directions[direction][1]

		outx := newx >= rows || newx < 0
		outy := newy >= columns || newy < 0
		if outx || outy || visited[newx][newy] == true {
			direction = (direction + 1) % len(directions)
		}

		x += directions[direction][0]
		y += directions[direction][1]
	}

	return ans
}