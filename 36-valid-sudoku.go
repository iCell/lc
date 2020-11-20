func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9, 9)
	columns := make([]map[byte]bool, 9, 9)
	grids := make([]map[byte]bool, 9, 9)
	
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		columns[i] = make(map[byte]bool, 9)
		grids[i] = make(map[byte]bool, 9)
	}
	
	for i := range board {
		for j, v := range board[i] {
			if v == '.' {
				continue
			}
			
			gridIdx := (i / 3) * 3 + j / 3
			
			if rows[i][v] || columns[j][v] || grids[gridIdx][v] {
				return false
			}
			
			rows[i][v] = true
			columns[j][v] = true
			grids[gridIdx][v] = true
		}
	}
	
	return true
}