

import (
    "mat    "sort"
h"
)

func minimumAbsDifference(arr []int) [][]int {
    if len(arr) == 0 {
        return nil
    }
    sort.Ints(arr)

    var result [][]int
    best := math.MaxInt64
    for i := 1; i < len(arr); i++ {
        sub := arr[i] - arr[i-1]
        if best > sub {
            best = sub
            result = [][]int{}
            result = append(result, []int{arr[i-1], arr[i]})
        } else if best == sub {
            result = append(result, []int{arr[i-1], arr[i]})
        }
    }

    return result
}

