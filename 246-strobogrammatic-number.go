var memo = map[byte]byte {
    '0': '0',
    '1': '1',
    '6': '9',
    '8': '8',
    '9': '6',
}

func isStrobogrammatic(num string) bool {
    low, high := 0, len(num) - 1
    for low < high {
        _, ok := memo[num[low]]
        if !ok {
            return false
        }
        _, ok = memo[num[high]]
        if !ok {
            return false
        }
        if memo[num[low]] != num[high] {
            return false
        }
        low++
        high--
    }
    if len(num) % 2 == 1 {
        v, ok := memo[num[low]]
        return ok && v == num[low]
    }
    return true
}
