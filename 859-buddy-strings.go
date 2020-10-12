func buddyStrings(A string, B string) bool {
    if len(A) != len(B) {
        return false
    }

    if A == B {
        memo := make(map[rune]int, 26)
        for _, b := range A {
            memo[b] += 1
        }
        var r bool
        for _, v := range memo {
            if v >= 2 {
                r = true
            }
        }
        return r
    }

    first, second := -1, -1
    for i := 0; i < len(A); i++ {
        if A[i] == B[i] {
            continue
        }
        if first != -1 {
            if second != -1 {
                return false
            }
            second = i
        } else {
            first = i
        }
    }
    return second != -1 && A[first] == B[second] && A[second] == B[first]
}