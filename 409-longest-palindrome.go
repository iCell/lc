func longestPalindrome(s string) int {
    hash := make(map[rune]int, len(s))
    for _, c := range s {
        hash[c] += 1
    }
    var length int
    var odd bool
    for _, v := range hash {
        length += v / 2 * 2
        if v%2 == 1 {
            odd = true
        }
    }
    if odd {
        length += 1
    }
    return length
}
