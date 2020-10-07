import "strings"

func numUniqueEmails(emails []string) int {
    memo := make(map[string]bool, len(emails))
    for _, email := range emails {
        memo[format(email)] = true
    }
    return len(memo)
}

func format(email string) string {
    components := strings.Split(email, "@")
    local, domain := components[0], components[1]

    local = (strings.Split(local, "+"))[0]
    local = strings.ReplaceAll(local, ".", "")

    return local + "@" + domain
}
