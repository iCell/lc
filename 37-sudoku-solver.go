func solveSudoku(board [][]byte)  {
	dfs(board, 0, 0)
}

func isValid(board [][]byte, row, col int, value byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == value {
			return false
		}
		if board[i][col] == value {
			return false
		}
		if board[(row/3)*3 + i/3][(col/3)*3 + i%3] == value {
			return false
		}
	}
	return true
}

func dfs(board [][]byte, row, col int) bool {
	if col == 9 {
		return dfs(board, row + 1, 0)
	}
	if row == 9 {
		return true
	}
	
	for i := row; i < 9; i++ {
		for j := col; j < 9; j++ {
			if board[i][j] != '.' {
				return dfs(board, i, j+1)
			}
			for c := byte('1'); c <= byte('9'); c++ {
				if !isValid(board, i, j, c) {
					continue
				}
				board[i][j] = c
				if dfs(board, i, j+1) {
					return true
				}
				board[i][j] = '.'
			}
			return false
		}
	}
	return false
}

   