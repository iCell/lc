package main

func countSubstrings(s string) int {
    m := len(s)
    if m == 0 {
        return 0
    }
    arr := make([][]int, m, m)
    for i := 0; i < m; i++ {
        row := make([]int, m, m)
        arr[i] = row
    }

    var count int
    for start := m - 1; start >= 0; start-- {
        for end := start; end < m; end++ {
            if start == end {
                arr[start][end] = 1
            } else if start+1 == end && s[start] == s[end] {
                arr[start][end] = 1
            } else if s[start] == s[end] && arr[start+1][end-1] == 1 {
                arr[start][end] = 1
            }
            if arr[start][end] == 1 {
                count += 1
            }
        }
    }

    return count
}
