package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 1; i < n; i++ {
		num := obstacleGrid[0][i]
		if num == 0 && obstacleGrid[0][i-1] == 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < m; i++ {
		num := obstacleGrid[i][0]
		if num == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
