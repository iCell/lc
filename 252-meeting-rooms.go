func canAttendMeetings(intervals [][]int) bool {
    if len(intervals) < 2 {
        return true
    }
    
    sort.SliceStable(intervals, func(left, right int) bool {
        return intervals[left][0] < intervals[right][0]
    })
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < intervals[i-1][1] {
            return false
        }
    }
    
    return true
}

