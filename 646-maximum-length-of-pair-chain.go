import "sort"

func findLongestChain(pairs [][]int) int {
    if len(pairs) == 0 {
        return 0
    }

    var temp = Pair(pairs)
    sort.Sort(temp)
    pairs = temp

    dp := make([]int, len(pairs), len(pairs))
    dp[0] = 1

    length := 1
    for i := 1; i < len(pairs); i++ {
        var longest int
        for j := 0; j < i; j++ {
            if pairs[i][0] > pairs[j][1] {
                longest = max(longest, dp[j])
            }
        }
        dp[i] = longest + 1
        length = max(length, dp[i])
    }

    return length
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

type Pair [][]int

func (p Pair) Len() int {
    return len(p)
}
func (p Pair) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p Pair) Less(i, j int) bool {
    return p[i][0] < p[j][0]
}