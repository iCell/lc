import "strings"

func numUniqueEmails(emails []string) int {
    memo := make(map[string]bool)
    for _, email := range emails {
        s := remove(email)
        if memo[s] {
            continue
        }
        memo[s] = true
    }
    return len(memo)
}

func remove(s string) string {
    domainIdx := strings.Index(s, "@")

    local := strings.ReplaceAll(s[:domainIdx], ".", "")

    plusIdx := strings.Index(local, "+")
    if plusIdx == -1 {
        return local + s[domainIdx:]
    }

    return local[:plusIdx] + s[domainIdx:]
}
