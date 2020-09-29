import "sort"

type Envelop [][]int

func (e Envelop) Len() int {
    return len(e)
}

func (e Envelop) Less(i, j int) bool {
    return e[i][0] < e[j][0]
}

func (e Envelop) Swap(i, j int) {
    e[i], e[j] = e[j], e[i]
}

func maxEnvelopes(envelopes [][]int) int {
    if len(envelopes) == 0 {
        return 0
    }

    sort.Sort(Envelop(envelopes))

    dp := make([]int, len(envelopes), len(envelopes))
    dp[0] = 1

    result := dp[0]
    for i := 1; i < len(dp); i++ {
        var m int
        for j := 0; j < i; j++ {
            if envelopes[i][1] > envelopes[j][1] && envelopes[i][0] > envelopes[j][0] {
                m = max(m, dp[j])
            }
        }
        dp[i] = m + 1
        result = max(result, dp[i])
    }

    return result
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

