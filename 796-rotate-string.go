func rotateString(A string, B string) bool {
    if len(A) != len(B) {
        return false
    }
    if len(A) == 0 || len(B) == 0 {
        return true
    }
    
    ra := []rune(A)
    for i := 0; i < len(ra); i++ {
        first := ra[0]
        ra = append(ra[1:], first)
        if string(ra) == B {
            return true
        }
    }
    
    return false
}

func rotateString(A string, B string) bool {
    if len(A) != len(B) {
        return false
    }
    if len(A) == 0 || len(B) == 0 {
        return true
    }
    return strings.Contains(A+A, B)
}