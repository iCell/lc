package main

import "math"

// f[0][0] = grid[0][0]
// f[0][i] = grid[0][i-1] + grid[0][i]
// f[i][0] = grid[i-1][0] + grid[i][0]
// f[m][n] = min(f[m][n-1], f[m-1][n]) + grid[m][n]
func minPathSum(grid [][]int) int {
    m, n := len(grid), 0
    if m == 0 {
        return 0
    }

    n = len(grid[0])
    if n == 0 {
        return 0
    }

    for i := 1; i < n; i++ {
        grid[0][i] += grid[0][i-1]
    }
    for i := 1; i < m; i++ {
        grid[i][0] += grid[i-1][0]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1])))
        }
    }

    return grid[m-1][n-1]
}
