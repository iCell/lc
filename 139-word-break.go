func wordBreak(s string, wordDict []string) bool {
    hash := make(map[string]bool, len(wordDict))
    for _, word := range wordDict {
        hash[word] = true
    }
    memo := make(map[string]bool)
    return dfs(s, hash, memo)
}

func dfs(s string, hash map[string]bool, memo map[string]bool) bool {
    if len(s) == 0 {
        return true
    }
    cached, ok := memo[s]
    if ok {
        return cached
    }

    for i := 1; i <= len(s); i++ {
        prefix := s[:i]
        if hash[prefix] && dfs(s[i:], hash, memo) {
            memo[s] = true
            return true
        }
    }

    memo[s] = false
    return false
}