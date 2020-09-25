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

func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]bool, len(wordDict))
    for _, word := range wordDict {
        dict[word] = true
    }
    memo := make(map[int]bool, len(s))
    return dfs(s, dict, 0, memo)
}

func dfs(s string, dict map[string]bool, idx int, memo map[int]bool) bool {
    if idx == len(s) {
        return true
    }

    r, calculated := memo[idx]
    if calculated {
        return r
    }

    for i := idx + 1; i <= len(s); i++ {
        prefix := s[idx:i]
        if dict[prefix] && dfs(s, dict, i, memo) {
            memo[idx] = true
            return true
        }
    }

    memo[idx] = false
    return false
}