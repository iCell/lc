import (
    "sort"
    "strings"
    "unicode"
)

func arrangeWords(text string) string {
    runes := []rune(text)
    runes[0] = unicode.ToLower(runes[0])
    origin := string(runes)

    substrings := strings.Split(origin, " ")
    sort.SliceStable(substrings, func(i, j int) bool {
        return len(substrings[i]) < len(substrings[j])
    })

    sorted := []rune(strings.Join(substrings, " "))
    sorted[0] = unicode.ToUpper(sorted[0])
    return string(sorted)
}