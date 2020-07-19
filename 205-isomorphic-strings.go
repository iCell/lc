func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    return helper(s, t) && helper(t, s)
}

func helper(s, t string) bool {
    sr, tr := []rune(s), []rune(t)

    temp := make(map[rune]rune, len(s))
    for i, c := range sr {
        tc := tr[i]
        mc, exist := temp[c]
        if exist {
            if tc != mc {
                return false
            }
            continue
        }
        temp[c] = tc
    }

    return true
}