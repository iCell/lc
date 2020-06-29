package main

// define start and end, to make a substring
// if start == end, then there is only one character, must be a palindromic string
// if start + 1 == end, only if character at start and end are the same
// f[start][end] = f[start+1][end-1]
//   0  1  2
// 0       t
// 1    s
// 2 f
// calculate f, then calculate s, then calculate t
// so, m -> m-1, n from 0 to n
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
                count++
            }
        }
    }

    return count
}
