import "strings"

func isPalindrome(s string) bool {
    runes := []rune(strings.ToLower(s))
    i, j := 0, len(runes)-1
    for i < j {
        if isLetter(runes[i]) == false {
            i++
            continue
        }
        if isLetter(runes[j]) == false {
            j--
            continue
        }

        if runes[i] != runes[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func isLetter(c rune) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}