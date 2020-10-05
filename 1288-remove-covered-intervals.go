import "sort"

type Interval [][]int

func (e Interval) Len() int {
    return len(e)
}

func (e Interval) Less(i, j int) bool {
    if e[i][0] < e[j][0] {
        return true
    } else if e[i][0] == e[j][0] {
        return e[i][1] > e[j][1]
    }
    return false
}

func (e Interval) Swap(i, j int) {
    e[i], e[j] = e[j], e[i]
}

func removeCoveredIntervals(intervals [][]int) int {
    if len(intervals) < 2 {
        return len(intervals)
    }

    sort.Sort(Interval(intervals))

    var count int
    pre := 0
    for pre < len(intervals) {
        cur := pre + 1
        for cur < len(intervals) && intervals[cur][1] >= intervals[cur][0] && intervals[cur][0] >= intervals[pre][0] && intervals[cur][1] <= intervals[pre][1] {
            cur++
            count++
        }
        pre = cur
    }

    return len(intervals) - count
}