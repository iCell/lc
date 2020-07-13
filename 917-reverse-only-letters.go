func reverseOnlyLetters(S string) string {
    runes := []rune(S)

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

        runes[i], runes[j] = runes[j], runes[i]
        i++
        j--
    }

    return string(runes)
}

func isLetter(c rune) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}