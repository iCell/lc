package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	num := 0
	for row, rows := range grid {
		for column, value := range rows {
			if value == '0' {
				continue
			}
			dfs(row, column, grid)
			num++
		}
	}

	return num
}

func dfs(row int, column int, grid [][]byte) {
	if row < 0 || column < 0 || row >= len(grid) || column >= len(grid[0]) || grid[row][column] == '0' {
		return
	}
	grid[row][column] = '0'
	dfs(row+1, column, grid)
	dfs(row, column+1, grid)
	dfs(row-1, column, grid)
	dfs(row, column-1, grid)
}
