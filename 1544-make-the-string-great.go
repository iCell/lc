func makeGood(s string) string {
    if len(s) == 0 {
        return s
    }

    for i := 1; i < len(s); i++ {
        left, right := s[i-1], s[i]
        if left != right && abs(left, right) == 32 {
            return makeGood(s[:i-1] + s[i+1:])
        }
    }

    return s
}

func abs(a, b byte) int {
    if a > b {
        return int(a) - int(b)
    }
    return int(b) - int(a)
}