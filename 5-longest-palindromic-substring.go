// if s[start] == s[end]:
// dp[start][end] = dp[start+1][end-1]
// dp[i][i] = 1

func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    
    dp := make([][]int, len(s), len(s))
    for i := range dp {
        dp[i] = make([]int, len(s), len(s))
    }
    
    var start, length int
    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < len(s); j++ {
            if s[i] == s[j] {
                if j - i < 3 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
            
            if dp[i][j] != 1 {
                continue
            }
            if j - i + 1 > length {
                length = j - i + 1
                start = i
            }
        }
    }
    
    return string(s[start:start+length])
}