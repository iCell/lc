package main

var directions = [][]int{
	[]int{1, 1},
	[]int{1, 0},
	[]int{0, 1},
	[]int{1, -1},
	[]int{0, -1},
	[]int{-1, 0},
	[]int{-1, 1},
	[]int{-1, -1},
}

func updateBoard(board [][]byte, click []int) [][]byte {

	row, column := click[0], click[1]

	if board[row][column] == 'M' {
		board[row][column] = 'X'
		return board
	}

	dfs(board, row, column)

	return board
}

func dfs(board [][]byte, row, column int) {
	numr, numc := len(board), len(board[0])
	if row < 0 || column < 0 || row >= numr || column >= numc || board[row][column] != 'E' {
		return
	}

	numm := 0
	for _, d := range directions {
		nextr, nextc := row+d[0], column+d[1]
		if nextr < numr && nextc < numc && nextc >= 0 && nextr >= 0 && board[nextr][nextc] == 'M' {
			numm++
		}
	}

	switch {
	case numm > 0:
		board[row][column] = byte(numm + int('0'))
	case numm == 0:
		board[row][column] = 'B'
		for _, d := range directions {
			nextr, nextc := row+d[0], column+d[1]
			dfs(board, nextr, nextc)
		}
	}
}
