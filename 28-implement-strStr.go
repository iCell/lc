package main

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }

    hr, nr := []rune(haystack), []rune(needle)
    for i, r := range hr {
        if r == nr[0] {
            if i+len(nr) > len(hr) {
                return -1
            }
            if string(hr[i:i+len(nr)]) == needle {
                return i
            }
        }
    }
    return -1
}
