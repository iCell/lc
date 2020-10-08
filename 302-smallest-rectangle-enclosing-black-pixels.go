var (
    top    int
    left   int
    bottom int
    right  int
)

func minArea(image [][]byte, x int, y int) int {
    m, n := len(image), 0
    if m == 0 {
        return 0
    }
    n = len(image[0])

    top, bottom, left, right = x, x, y, y
    dfs(image, x, y, m, n)

    return (right - left) * (bottom - top)
}

func dfs(image [][]byte, x, y int, m, n int) {
    if x >= m || x < 0 || y >= n || y < 0 {
        return
    }
    if image[x][y] == '0' {
        return
    }
    image[x][y] = '0'

    top, bottom = min(top, x), max(bottom, x+1)
    left, right = min(left, y), max(right, y+1)

    dfs(image, x+1, y, m, n)
    dfs(image, x-1, y, m, n)
    dfs(image, x, y+1, m, n)
    dfs(image, x, y-1, m, n)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}