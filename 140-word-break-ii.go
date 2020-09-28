func wordBreak(s string, wordDict []string) []string {
    dict := make(map[string]bool, len(wordDict))
    for _, word := range wordDict {
        dict[word] = true
    }
    memo := make(map[string][]string)

    return helper(s, dict, memo)
}

func helper(s string, dict map[string]bool, memo map[string][]string) []string {
    var result []string
    if len(s) == 0 {
        result = append(result, "")
        return result
    }

    temp, calculated := memo[s]
    if calculated {
        return temp
    }

    for i := 1; i <= len(s); i++ {
        prefix := s[:i]
        _, contain := dict[prefix]
        if !contain {
            continue
        }
        remains := helper(s[i:], dict, memo)
        for _, remain := range remains {
            var option = " "
            if len(remain) == 0 {
                option = ""
            }
            result = append(result, prefix+option+remain)
        }
    }

    memo[s] = result

    return result
}