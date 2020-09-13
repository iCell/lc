var directions = [][]int{
    []int{1, 0},
    []int{0, 1},
    []int{0, -1},
    []int{-1, 0},
}

func exist(board [][]byte, word string) bool {
    bytes := []byte(word)
    m, n := len(board), len(board[0])
    visited := make([][]bool, m, m)
    for i := range visited {
        visited[i] = make([]bool, n, n)
    }
    for i := range board {
        for j := range board[i] {
            result := dfs(board, i, j, 0, bytes, visited)
            if result {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, x, y int, idx int, word []byte, visited [][]bool) bool {
    if board[x][y] != word[idx] {
        return false
    }
    if idx == len(word)-1 {
        return true
    }

    visited[x][y] = true
    for _, direction := range directions {
        newX, newY := x+direction[0], y+direction[1]
        if newX >= len(board) || newX < 0 || newY >= len(board[0]) || newY < 0 {
            continue
        }
        if visited[newX][newY] {
            continue
        }
        result := dfs(board, newX, newY, idx+1, word, visited)

        if result {
            return true
        }
    }
    visited[x][y] = false
    return false
}

