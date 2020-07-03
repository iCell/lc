func lengthOfLastWord(s string) int {
    var length int
    found := false
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            found = true
        }
        if s[i] == ' ' && found {
            break
        }
        if found {
            length++
        }
    }
    return length
}