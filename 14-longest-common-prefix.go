func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for idx, char := range strs[0] {
        for i := 1; i < len(strs); i++ {
            str := []rune(strs[i])
            if idx >= len(str) || char != str[idx] {
                return string(str[:idx])
            }
        }
    }
    return strs[0]
}

