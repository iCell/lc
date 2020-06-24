package main

import "fmt"

func main() {
    fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
    cache := make([][]int, m)
    for i := 0; i < m; i++ {
        cache[i] = make([]int, n)
        cache[i][0] = 1
    }
    for i := 0; i < n; i++ {
        cache[0][i] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            cache[i][j] = cache[i-1][j] + cache[i][j-1]
        }
    }
    return cache[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
    if m == 1 || n == 1 {
        return 1
    }

    cache := make([][]int, m, m)
    for i := 0; i < m; i++ {
        cache[i] = make([]int, n, n)
    }

    for i := 0; i < m; i++ {
        cache[i][0] = 1
    }
    for i := 0; i < n; i++ {
        cache[0][i] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            cache[i][j] = cache[i-1][j] + cache[i][j-1]
        }
    }
    return cache[m-1][n-1]
}
