func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	
	x, y := 0, n - 1
	for x < m && y >= 0 {
		value := matrix[x][y]
		if value == target {
			return true
		} else if value > target {
			y--
		} else if target > value {
			x++
		}
	}
	return false
}
