func fourSumCount(A []int, B []int, C []int, D []int) int {
    cache := make(map[int]int)
    for _, a := range A {
        for _, b := range B {
            cache[a+b] += 1
        }
    }

    var count int
    for _, c := range C {
        for _, d := range D {
            key := -(c + d)
            v, ok := cache[key]
            if ok {
                count += v
            }
        }
    }

    return count
}